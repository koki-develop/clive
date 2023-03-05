package ui

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"github.com/koki-develop/clive/pkg/browser"
	"github.com/koki-develop/clive/pkg/config"
	"github.com/koki-develop/clive/pkg/ttyd"
	"github.com/koki-develop/clive/pkg/util"
)

type loadConfigMsg struct{ config *config.Config }
type startTtydMsg struct{ ttyd *ttyd.Ttyd }
type openMsg struct{ page *rod.Page }
type runMsg struct{}
type pauseMsg struct{}
type quitMsg struct{}
type errMsg struct{ err error }

// $ base64 ./assets/favicon.png
var favicon = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAACXBIWXMAAAsTAAALEwEAmpwYAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAACASURBVHgB7dRLCsAgDEXRt7TsPEuzCi2IqIVic+nnQebnTiI9fUnJRC4DXNRKfb4kaqUeAxz1JMAxQF1PARwDtPUEwDFArz4a4BhgVD+5tS96VD8506r99d+tP4HdX38BYIoaWj8AmCKH1ncApuih9Q3ARAytrwAmamj9DjC9YRvd3NI8H6lTHgAAAABJRU5ErkJggg=="

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// spinner
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	// key
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, m.quit
		case tea.KeyEnter:
			if m.quitting {
				return m, tea.Quit
			}
			if m.pausing {
				m.pausing = false
				m.currentActionIndex++
				return m, m.run
			}
		}

	// events
	case loadConfigMsg:
		m.config = msg.config
		return m, m.startTtyd
	case startTtydMsg:
		m.ttyd = msg.ttyd
		return m, m.open
	case openMsg:
		m.page = msg.page
		return m, tea.Batch(tea.EnterAltScreen, m.run)
	case runMsg:
		m.currentActionIndex++
		return m, m.run
	case pauseMsg:
		m.pausing = true
		return m, nil
	case quitMsg:
		m.quitting = true
		if m.config.Settings.SkipPauseBeforeQuit {
			return m, tea.Quit
		}
		return m, nil
	case errMsg:
		m.err = msg.err
		if !m.running() {
			return m, tea.Quit
		}
		return m, m.quit
	}

	return m, nil
}

func (m *Model) loadConfig() tea.Msg {
	cfg, err := config.Load(m.configFile)
	if err != nil {
		return errMsg{err}
	}

	return loadConfigMsg{cfg}
}

func (m *Model) startTtyd() tea.Msg {
	port, err := m.netListener.RandomUnusedTCPPort()
	if err != nil {
		return errMsg{err}
	}

	ttyd := ttyd.New(m.config.Settings.LoginCommand, port)
	if err := ttyd.Start(); err != nil {
		return errMsg{err}
	}

	return startTtydMsg{ttyd}
}

func (m *Model) open() tea.Msg {
	page, err := m.openPage()
	if err != nil {
		return errMsg{err}
	}

	return openMsg{page}
}

func (m *Model) openPage() (*rod.Page, error) {
	url := fmt.Sprintf("http://localhost:%d", m.ttyd.Port)
	p, err := browser.Open(&browser.BrowserConfig{Bin: m.config.Settings.BrowserBin, URL: url, Headless: m.config.Settings.Headless})
	if err != nil {
		return nil, err
	}

	if err := m.setupPage(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (m *Model) setupPage(page *rod.Page) error {
	// window size
	if m.config.Settings.Width != nil || m.config.Settings.Height != nil {
		if err := page.SetWindow(&proto.BrowserBounds{Width: m.config.Settings.Width, Height: m.config.Settings.Height}); err != nil {
			return err
		}
	}

	// favicon
	if _, err := page.Eval(fmt.Sprintf(`() => document.querySelector('link[rel="icon"]').href = "%s"
`, favicon)); err != nil {
		return err
	}

	// font family
	if m.config.Settings.FontFamily != nil {
		if _, err := page.Eval(fmt.Sprintf("() => term.options.fontFamily = '%s'", *m.config.Settings.FontFamily)); err != nil {
			return err
		}
	}

	// font size
	if _, err := page.Eval(fmt.Sprintf("() => term.options.fontSize = %d", m.config.Settings.FontSize)); err != nil {
		return err
	}
	if _, err := page.Eval("term.fit"); err != nil {
		return err
	}

	return nil
}
func (m *Model) quit() tea.Msg {
	return quitMsg{}
}

func (m *Model) run() tea.Msg {
	if m.currentActionIndex == len(m.config.Actions) {
		return m.quit()
	}

	action := m.config.Actions[m.currentActionIndex]
	switch action := action.(type) {
	case *config.PauseAction:
		return m.runPause(action)
	case *config.TypeAction:
		return m.runType(action)
	case *config.KeyAction:
		return m.runKey(action)
	case *config.SleepAction:
		return m.runSleep(action)
	case *config.CtrlAction:
		return m.runCtrl(action)
	case *config.ScreenshotAction:
		return m.runScreenshot(action)
	default:
		return errMsg{fmt.Errorf("unknown action: %#v", action)}
	}
}

func (m *Model) runPause(action *config.PauseAction) tea.Msg {
	return pauseMsg{}
}

func (m *Model) runType(action *config.TypeAction) tea.Msg {
	for i := 0; i < action.Count; i++ {
		if err := m.runTypeOnce(action); err != nil {
			return errMsg{err}
		}
		if m.quitting {
			return nil
		}
	}

	return runMsg{}
}

func (m *Model) runTypeOnce(action *config.TypeAction) error {
	for _, c := range action.Type {
		if err := m.typeChar(c); err != nil {
			return err
		}
		time.Sleep(time.Duration(action.Speed) * time.Millisecond)
		if m.quitting {
			return nil
		}
	}

	return nil
}

func (m *Model) typeChar(c rune) error {
	k, ok := config.KeyMap[c]
	if ok {
		if err := m.page.Keyboard.Type(k); err != nil {
			return err
		}
	} else {
		area, err := m.page.Element("textarea")
		if err != nil {
			return err
		}
		if err := m.input(area, string(c)); err != nil {
			return err
		}
	}

	return nil
}

func (m *Model) input(elm *rod.Element, text string) error {
	if err := elm.Input(text); err != nil {
		return err
	}

	if err := m.page.WaitIdle(time.Minute); err != nil {
		return err
	}

	return nil
}

func (m *Model) runKey(action *config.KeyAction) tea.Msg {
	for i := 0; i < action.Count; i++ {
		if err := m.runKeyOnce(action); err != nil {
			return errMsg{err}
		}
		if m.quitting {
			return nil
		}
	}

	return runMsg{}
}

func (m *Model) runKeyOnce(action *config.KeyAction) error {
	k, ok := config.SpecialKeyMap[action.Key]
	if !ok {
		return nil
	}

	if err := m.page.Keyboard.Type(k); err != nil {
		return err
	}

	time.Sleep(time.Duration(action.Speed) * time.Millisecond)
	return nil
}

func (m *Model) runSleep(action *config.SleepAction) tea.Msg {
	time.Sleep(time.Duration(action.Sleep) * time.Millisecond)
	if m.quitting {
		return nil
	}
	return runMsg{}
}

func (m *Model) runCtrl(action *config.CtrlAction) tea.Msg {
	for i := 0; i < action.Count; i++ {
		if err := m.runCtrlOnce(action); err != nil {
			return errMsg{err}
		}
		if m.quitting {
			return nil
		}
	}

	return runMsg{}
}

func (m *Model) runCtrlOnce(action *config.CtrlAction) error {
	if err := m.page.Keyboard.Press(input.ControlLeft); err != nil {
		return err
	}
	for _, r := range action.Ctrl {
		k, ok := config.KeyMap[r]
		if !ok {
			continue
		}
		if err := m.page.Keyboard.Type(k); err != nil {
			return err
		}
	}
	if err := m.page.Keyboard.Release(input.ControlLeft); err != nil {
		return err
	}

	time.Sleep(time.Duration(action.Speed) * time.Millisecond)
	return nil
}

func (m *Model) runScreenshot(action *config.ScreenshotAction) tea.Msg {
	if err := m.setCursorBlink(false); err != nil {
		return errMsg{err}
	}

	buf, err := m.page.Screenshot(false, nil)
	if err != nil {
		return errMsg{err}
	}

	if err := m.setCursorBlink(true); err != nil {
		return errMsg{err}
	}

	filename := fmt.Sprintf("%d_%s.png", m.currentActionIndex+1, time.Now().Format("20060102150405"))
	if action.File != nil {
		filename = *action.File
	}
	p := filepath.Join(action.Dir, filename)
	f, err := util.CreateFile(p)
	if err != nil {
		return errMsg{err}
	}
	defer f.Close()

	if _, err := f.Write(buf); err != nil {
		return errMsg{err}
	}

	return runMsg{}
}

func (m *Model) setCursorBlink(v bool) error {
	if _, err := m.page.Eval(fmt.Sprintf("() => term.options.cursorBlink = %t", v)); err != nil {
		return err
	}
	return nil
}
