package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
	"github.com/koki-develop/clive/pkg/config"
	"github.com/koki-develop/clive/pkg/ttyd"
	"github.com/koki-develop/clive/pkg/ui"
	"github.com/koki-develop/clive/pkg/util"
	"github.com/spf13/cobra"
)

var (
	actionsHeaderStyle = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(lipgloss.Color("#ff00ff"))
	errHeaderStyle     = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(lipgloss.Color("#ff0000"))
)

type startModel struct {
	Err                error
	Spinner            spinner.Model
	Config             *config.Config
	Ttyd               *ttyd.Ttyd
	Browser            *rod.Browser
	Page               *rod.Page
	CurrentActionIndex int
	Running            bool
	Pausing            bool
	PausingBeforeQuit  bool
}

func newStartModel() *startModel {
	s := spinner.New(
		spinner.WithSpinner(spinner.Dot),
		spinner.WithStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff00ff"))),
	)

	return &startModel{
		Spinner:            s,
		CurrentActionIndex: 0,
	}
}

func (m *startModel) loadConfig() tea.Msg {
	cfg, err := config.Load(configFilename)
	if err != nil {
		return errMsg{err}
	}

	return configLoadedMsg{cfg}
}

func (m *startModel) startTtyd() tea.Msg {
	ttyd, err := ttyd.NewTtyd(m.Config.Settings.LoginCommand)
	if err != nil {
		return errMsg{err}
	}

	if err := ttyd.Command.Start(); err != nil {
		return errMsg{err}
	}

	return ttydStartedMsg{ttyd}
}

func (m *startModel) launchBrowser() tea.Msg {
	browser, err := launchBrowser(m.Config)
	if err != nil {
		return errMsg{err}
	}

	return browserLaunchedMsg{browser}
}

func (m *startModel) openPage() tea.Msg {
	page, err := m.Browser.Page(proto.TargetCreateTarget{URL: fmt.Sprintf("http://localhost:%d", m.Ttyd.Port)})
	if err != nil {
		return errMsg{err}
	}
	if err := page.WaitIdle(time.Minute); err != nil {
		return errMsg{err}
	}
	if err := m.setupPage(page); err != nil {
		return errMsg{err}
	}

	return pageOpenedMsg{page}
}

func (m *startModel) setupPage(page *rod.Page) error {
	if m.Config.Settings.FontFamily != nil {
		if _, err := page.Eval(fmt.Sprintf("() => term.options.fontFamily = '%s'", *m.Config.Settings.FontFamily)); err != nil {
			return err
		}
	}
	if _, err := page.Eval(fmt.Sprintf("() => term.options.fontSize = %d", m.Config.Settings.FontSize)); err != nil {
		return err
	}
	if _, err := page.Eval("term.fit"); err != nil {
		return err
	}

	return nil
}

func (m *startModel) runAction() tea.Msg {
	if m.CurrentActionIndex == len(m.Config.Actions) {
		return pauseBeforeQuitMsg{}
	}

	action := m.Config.Actions[m.CurrentActionIndex]

	switch action := action.(type) {
	case *config.PauseAction:
		return m.runPauseAction(action)
	case *config.TypeAction:
		return m.runTypeAction(action)
	case *config.KeyAction:
		return m.runKeyAction(action)
	case *config.SleepAction:
		return m.runSleepAction(action)
	case *config.CtrlAction:
		return m.runCtrlAction(action)
	default:
		return errMsg{fmt.Errorf("unknown action: %#v", action)}
	}
}

func (m *startModel) runPauseAction(action *config.PauseAction) tea.Msg {
	return pauseActionMsg{}
}

func (m *startModel) runTypeAction(action *config.TypeAction) tea.Msg {
	for i := 0; i < action.Count; i++ {
		for _, c := range action.Type {
			k, ok := config.KeyMap[c]
			if ok {
				if err := m.Page.Keyboard.Type(k); err != nil {
					return errMsg{err}
				}
			} else {
				txt, err := m.Page.Element("textarea")
				if err != nil {
					return errMsg{err}
				}
				if err := txt.Input(string(c)); err != nil {
					return errMsg{err}
				}
				if err := m.Page.WaitIdle(time.Minute); err != nil {
					return errMsg{err}
				}
			}
			time.Sleep(time.Duration(action.Speed) * time.Millisecond)
			if m.PausingBeforeQuit {
				return nil
			}
		}
	}
	return actionDoneMsg{}
}

func (m *startModel) runKeyAction(action *config.KeyAction) tea.Msg {
	k, ok := config.SpecialKeyMap[action.Key]
	for i := 0; i < action.Count; i++ {
		if ok {
			if err := m.Page.Keyboard.Type(k); err != nil {
				return errMsg{err}
			}
		}
		time.Sleep(time.Duration(action.Speed) * time.Millisecond)
		if m.PausingBeforeQuit {
			return nil
		}
	}
	return actionDoneMsg{}
}

func (m *startModel) runSleepAction(action *config.SleepAction) tea.Msg {
	time.Sleep(time.Duration(action.Sleep) * time.Millisecond)
	if m.PausingBeforeQuit {
		return nil
	}
	return actionDoneMsg{}
}

func (m *startModel) runCtrlAction(action *config.CtrlAction) tea.Msg {
	for i := 0; i < action.Count; i++ {
		if err := m.Page.Keyboard.Press(input.ControlLeft); err != nil {
			return errMsg{err}
		}
		for _, r := range action.Ctrl {
			if k, ok := config.KeyMap[r]; ok {
				if err := m.Page.Keyboard.Type(k); err != nil {
					return errMsg{err}
				}
			}
		}
		if err := m.Page.Keyboard.Release(input.ControlLeft); err != nil {
			return errMsg{err}
		}
		time.Sleep(time.Duration(action.Speed) * time.Millisecond)
		if m.PausingBeforeQuit {
			return nil
		}
	}
	return actionDoneMsg{}
}

func (m *startModel) Init() tea.Cmd {
	return tea.Batch(
		m.Spinner.Tick,
		m.loadConfig,
	)
}

func (m *startModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			m.PausingBeforeQuit = true
			return m, nil
		case tea.KeyEnter:
			if m.PausingBeforeQuit {
				return m, tea.Quit
			}
			if m.Pausing {
				m.Pausing = false
				m.CurrentActionIndex++
				return m, m.runAction
			}
		}
	case configLoadedMsg:
		m.Config = msg.config
		return m, m.startTtyd
	case ttydStartedMsg:
		m.Ttyd = msg.Ttyd
		return m, m.launchBrowser
	case browserLaunchedMsg:
		m.Browser = msg.Browser
		return m, m.openPage
	case pageOpenedMsg:
		m.Page = msg.Page
		m.Running = true
		return m, tea.Batch(tea.EnterAltScreen, m.runAction)
	case pauseActionMsg:
		m.Pausing = true
		return m, nil
	case pauseBeforeQuitMsg:
		m.PausingBeforeQuit = true
		return m, nil
	case actionDoneMsg:
		m.CurrentActionIndex++
		return m, m.runAction
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	case errMsg:
		m.Err = msg.Err
		if !m.Running {
			return m, tea.Quit
		}

		m.PausingBeforeQuit = true
		return m, nil
	}

	return m, nil
}

func (m *startModel) errorView() string {
	if m.Err == nil {
		return ""
	}
	return errHeaderStyle.Render("Error") + "\n" + fmt.Sprint(m.Err)
}

func (m *startModel) loadingConfigView() string {
	return fmt.Sprintf("%s Loading config", m.Spinner.View())
}

func (m *startModel) launchingBrowserView() string {
	return fmt.Sprintf("%s Launching browser", m.Spinner.View())
}

func (m *startModel) pauseBeforeQuitView() string {
	return lipgloss.NewStyle().Bold(true).Render("Press enter to quit")
}

func (m *startModel) actionsView() string {
	from := util.Max(0, m.CurrentActionIndex-3)
	show := 20
	digits := len(strconv.Itoa(len(m.Config.Actions)))

	rows := []string{}
	for i, action := range m.Config.Actions {
		if i < from && len(m.Config.Actions)-i > show {
			continue
		}
		if i-from >= show {
			rows = append(rows, fmt.Sprintf("... %d more actions", len(m.Config.Actions)-i))
			break
		}

		style := lipgloss.NewStyle()

		cursor := "  "
		if m.CurrentActionIndex > i {
			style = style.Faint(true)
		} else if m.CurrentActionIndex == i {
			style = style.Bold(true)
			if !m.PausingBeforeQuit {
				if m.Pausing {
					cursor = "> "
				} else {
					cursor = m.Spinner.View()
				}
			}
		}

		num := util.PaddingRight(fmt.Sprintf("#%d", i+1), digits+1)
		rows = append(rows, fmt.Sprintf("%s %s%s", style.Render(num), cursor, style.Render(action.String())))
	}

	return actionsHeaderStyle.Render("Actions") + "\n" + strings.Join(rows, "\n")
}

func (m *startModel) View() string {
	s := ""

	if m.Err != nil {
		if !m.Running {
			return ""
		}

		s += m.errorView()
		s += "\n\n" + m.pauseBeforeQuitView()
		return s
	}

	if m.Config == nil {
		s += m.loadingConfigView()
		return s
	}

	if m.Browser == nil || m.Page == nil {
		s += m.launchingBrowserView()
		return s
	}

	s += m.actionsView()

	if m.PausingBeforeQuit {
		s += "\n\n" + lipgloss.NewStyle().Bold(true).Render("Press enter to quit")
	}

	return s
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start cLive actions",
	Long:  "Start cLive actions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		m := ui.New(configFilename)
		defer m.Close()

		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			return err
		}

		if err := m.Err(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
