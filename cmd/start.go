package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/spf13/cobra"
)

type startModel struct {
	err                error
	spinner            spinner.Model
	config             *config
	port               int
	ttyd               *exec.Cmd
	browser            *rod.Browser
	page               *rod.Page
	currentActionIndex int
	pausing            bool
}

func newStartModel() *startModel {
	s := spinner.New(
		spinner.WithSpinner(spinner.Dot),
		spinner.WithStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff00ff"))),
	)

	return &startModel{
		spinner:            s,
		currentActionIndex: 0,
	}
}

type configLoadedMsg struct {
	config *config
}

type ttydStartedMsg struct {
	port int
	ttyd *exec.Cmd
}

type browserLaunchedMsg struct {
	browser *rod.Browser
	page    *rod.Page
}

type actionDoneMsg struct{}

type pauseActionMsg struct{}

type errMsg struct {
	err error
}

func (m *startModel) loadConfig() tea.Msg {
	cfg, err := loadConfig(configFilename)
	if err != nil {
		return errMsg{err}
	}

	return configLoadedMsg{cfg}
}

func (m *startModel) startTtyd() tea.Msg {
	port, err := randomUnusedPort()
	if err != nil {
		return errMsg{err}
	}

	ttyd := ttyd(port)
	if err := ttyd.Start(); err != nil {
		return errMsg{err}
	}

	return ttydStartedMsg{port, ttyd}
}

func (m *startModel) launchBrowser() tea.Msg {
	browser, err := launchBrowser()
	if err != nil {
		return errMsg{err}
	}

	page := browser.
		NoDefaultDevice().
		MustPage(fmt.Sprintf("http://localhost:%d", m.port)).
		MustWaitIdle()
	_ = page.MustEval("() => term.options.fontSize = 22")
	_ = page.MustEval("term.fit")

	return browserLaunchedMsg{browser, page}
}

func (m *startModel) runAction() tea.Msg {
	action := m.config.Actions[m.currentActionIndex]

	switch action := action.(type) {
	case *pauseAction:
		return pauseActionMsg{}
	case *typeAction:
		for _, c := range action.Type {
			k, ok := keymap[c]
			if ok {
				_ = m.page.Keyboard.MustType(k)
			} else {
				_ = m.page.MustElement("textarea").Input(string(c))
				_ = m.page.MustWaitIdle()
			}
			time.Sleep(time.Duration(action.Speed) * time.Millisecond)
		}
	case *keyAction:
		k, ok := specialkeymap[strings.ToLower(action.Key)]
		for i := 0; i < action.Count; i++ {
			if ok {
				_ = m.page.Keyboard.MustType(k)
			}
			time.Sleep(time.Duration(action.Speed) * time.Millisecond)
		}
	case *sleepAction:
		time.Sleep(time.Duration(action.Time) * time.Millisecond)
	case *ctrlAction:
		_ = m.page.Keyboard.Press(input.ControlLeft)
		for _, r := range action.Ctrl {
			if k, ok := keymap[r]; ok {
				_ = m.page.Keyboard.Type(k)
			}
		}
		_ = m.page.Keyboard.Release(input.ControlLeft)
	}
	return actionDoneMsg{}
}

func (m *startModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.loadConfig,
	)
}

func (m *startModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.pausing {
				m.pausing = false
				m.currentActionIndex++
				if m.currentActionIndex == len(m.config.Actions) {
					return m, tea.Quit
				}
				return m, m.runAction
			}
		}
	case configLoadedMsg:
		m.config = msg.config
		return m, m.startTtyd
	case ttydStartedMsg:
		m.ttyd = msg.ttyd
		m.port = msg.port
		return m, m.launchBrowser
	case browserLaunchedMsg:
		m.browser = msg.browser
		m.page = msg.page
		return m, tea.Batch(tea.EnterAltScreen, m.runAction)
	case pauseActionMsg:
		m.pausing = true
		return m, nil
	case actionDoneMsg:
		m.currentActionIndex++
		if m.currentActionIndex == len(m.config.Actions) {
			return m, tea.Quit
		}
		return m, m.runAction
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case errMsg:
		m.err = msg.err
		return m, tea.Quit
	}

	return m, nil
}

func (m *startModel) View() string {
	if m.err != nil {
		return ""
	}

	if m.config == nil {
		return fmt.Sprintf("%s Loading config", m.spinner.View())
	}

	if m.browser == nil {
		return fmt.Sprintf("%s Launching browser", m.spinner.View())
	}

	s := ""

	from := max(0, m.currentActionIndex-8)

	for i, action := range m.config.Actions {
		if i < from {
			continue
		}
		if i-from >= 20 {
			s += fmt.Sprintf("... %d more actions", len(m.config.Actions)-i)
			break
		}

		cursor := "  "
		text := action.String()

		if m.currentActionIndex > i {
			text = color.New(color.Faint).Sprint(text)
		} else if m.currentActionIndex == i {
			text = color.New(color.Bold).Sprint(text)
			if m.pausing {
				cursor = "> "
			} else {
				cursor = m.spinner.View()
			}
		}
		s += fmt.Sprintf("%s%s\n", cursor, text)
	}

	return s
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start clive actions",
	Long:  "Start clive actions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		m := newStartModel()
		defer func() {
			if m.ttyd != nil {
				_ = m.ttyd.Process.Kill()
			}
		}()

		p := tea.NewProgram(m)
		if err := p.Start(); err != nil {
			return err
		}

		if m.err != nil {
			return m.err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
