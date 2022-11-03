package cmd

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
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
	pausingBeforeQuit  bool
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

type pauseBeforeQuitMsg struct{}

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

	ttyd := ttyd(port, m.config.Settings.LoginCommand)
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

	page, err := browser.Page(proto.TargetCreateTarget{URL: fmt.Sprintf("http://localhost:%d", m.port)})
	if err != nil {
		return errMsg{err}
	}
	if err := page.WaitIdle(time.Minute); err != nil {
		return errMsg{err}
	}

	if m.config.Settings.FontFamily != nil {
		if _, err := page.Eval(fmt.Sprintf("() => term.options.fontFamily = '%s'", *m.config.Settings.FontFamily)); err != nil {
			return errMsg{err}
		}
	}
	if _, err = page.Eval(fmt.Sprintf("() => term.options.fontSize = %d", m.config.Settings.FontSize)); err != nil {
		return errMsg{err}
	}
	if _, err = page.Eval("term.fit"); err != nil {
		return errMsg{err}
	}

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
		time.Sleep(time.Duration(action.Sleep) * time.Millisecond)
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

func (m *startModel) pauseBeforeQuit() tea.Msg {
	return pauseBeforeQuitMsg{}
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
					return m, m.pauseBeforeQuit
				}
				return m, m.runAction
			}
			if m.pausingBeforeQuit {
				return m, tea.Quit
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
	case pauseBeforeQuitMsg:
		m.pausingBeforeQuit = true
		return m, nil
	case actionDoneMsg:
		m.currentActionIndex++
		if m.currentActionIndex == len(m.config.Actions) {
			return m, m.pauseBeforeQuit
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

	s := lipgloss.NewStyle().Bold(true).Background(lipgloss.Color("#ff00ff")).Padding(0, 1).Render("Actions") + "\n"

	from := max(0, m.currentActionIndex-3)
	show := 20
	digits := len(strconv.Itoa(len(m.config.Actions)))

	for i, action := range m.config.Actions {
		if i < from && len(m.config.Actions)-i > show {
			continue
		}
		if i-from >= show {
			s += fmt.Sprintf("... %d more actions", len(m.config.Actions)-i)
			break
		}

		style := lipgloss.NewStyle()

		cursor := "  "
		if m.currentActionIndex > i {
			style = style.Faint(true)
		} else if m.currentActionIndex == i {
			style = style.Bold(true)
			if m.pausing {
				cursor = "> "
			} else {
				cursor = m.spinner.View()
			}
		}

		num := paddingRight(fmt.Sprintf("#%d", i+1), digits+1)
		s += fmt.Sprintf("%s %s%s\n", style.Render(num), cursor, style.Render(action.String()))
	}

	if m.pausingBeforeQuit {
		s += "\n" + lipgloss.NewStyle().Bold(true).Render("Press enter to quit")
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
