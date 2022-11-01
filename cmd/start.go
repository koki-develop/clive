package cmd

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type startModel struct {
	config             *config
	currentActionIndex int
	spinner            spinner.Model
}

func newStartModel(cfg *config) startModel {
	s := spinner.New(
		spinner.WithSpinner(spinner.Dot),
		spinner.WithStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff00ff"))),
	)

	return startModel{
		config:             cfg,
		currentActionIndex: 0,
		spinner:            s,
	}
}

func (m startModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m startModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			m.currentActionIndex++
			if m.currentActionIndex == len(m.config.Actions) {
				return m, tea.Quit
			}
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m startModel) View() string {
	s := ""
	for i, action := range m.config.Actions {
		cursor := "  "
		if m.currentActionIndex == i {
			cursor = m.spinner.View()
		}
		s += fmt.Sprintf("%s%s\n", cursor, action)
	}

	return s
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start clive actions",
	Long:  "Start clive actions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgname, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		cfg, err := loadConfig(cfgname)
		if err != nil {
			return err
		}

		p := tea.NewProgram(newStartModel(cfg))
		if err := p.Start(); err != nil {
			return err
		}

		return nil

		// s := spinner.New(spinner.CharSets[11], 100*time.Millisecond, spinner.WithWriter(os.Stderr))
		// if err := s.Color("magenta"); err != nil {
		// 	return err
		// }
		// defer s.Stop()
		// ps := spinner.New([]string{">"}, 100*time.Millisecond, spinner.WithWriter(os.Stderr))
		// if err := ps.Color("magenta"); err != nil {
		// 	return err
		// }
		// defer ps.Stop()

		// port, err := randomUnusedPort()
		// if err != nil {
		// 	return err
		// }

		// s.Suffix = " Starting ttyd"
		// s.Start()
		// ttyd := ttyd(port)
		// if err := ttyd.Start(); err != nil {
		// 	return err
		// }
		// defer ttyd.Process.Kill()
		// s.Stop()

		// s.Suffix = " Launching browser"
		// s.Start()
		// browser, err := launchBrowser()
		// if err != nil {
		// 	return err
		// }
		// s.Stop()

		// s.Suffix = " Opening page"
		// s.Start()
		// page := browser.
		// 	NoDefaultDevice().
		// 	MustPage(fmt.Sprintf("http://localhost:%d", port)).
		// 	MustWaitIdle()
		// _ = page.MustEval("() => term.options.fontSize = 22")
		// _ = page.MustEval("term.fit")
		// s.Stop()

		// for i, action := range cfg.Actions {
		// 	switch action := action.(type) {
		// 	case *typeAction:
		// 		s.Suffix = " " + color.New(color.Bold).Sprint(action)
		// 		s.Start()

		// 		for _, c := range action.Type {
		// 			k, ok := keymap[c]
		// 			if ok {
		// 				_ = page.Keyboard.MustType(k)
		// 			} else {
		// 				_ = page.MustElement("textarea").Input(string(c))
		// 				_ = page.MustWaitIdle()
		// 			}
		// 			time.Sleep(time.Duration(action.Speed) * time.Millisecond)
		// 		}

		// 		s.Stop()
		// 		fmt.Println(action)
		// 	case *keyAction:
		// 		s.Suffix = " " + color.New(color.Bold).Sprint(action)
		// 		s.Start()

		// 		k, ok := specialkeymap[strings.ToLower(action.Key)]
		// 		for i := 0; i < action.Count; i++ {
		// 			if ok {
		// 				_ = page.Keyboard.MustType(k)
		// 			}
		// 			time.Sleep(time.Duration(action.Speed) * time.Millisecond)
		// 		}

		// 		s.Stop()
		// 		fmt.Println(action)
		// 	case *pauseAction:
		// 		next := "quit"
		// 		if i+1 < len(cfg.Actions) {
		// 			next = cfg.Actions[i+1].String()
		// 		}
		// 		log := fmt.Sprintf("%s (Next: %s)", color.New(color.Bold).Sprint(action), next)
		// 		ps.Suffix = " " + log
		// 		ps.Start()

		// 		for {
		// 			_, key, err := keyboard.GetSingleKey()
		// 			if err != nil {
		// 				return err
		// 			}
		// 			if key == keyboard.KeyEnter {
		// 				break
		// 			}
		// 		}

		// 		ps.Stop()
		// 	case *sleepAction:
		// 		s.Suffix = " " + color.New(color.Bold).Sprint(action)
		// 		s.Start()

		// 		time.Sleep(time.Duration(action.Time) * time.Millisecond)

		// 		s.Stop()
		// 		fmt.Println(action)
		// 	case *ctrlAction:
		// 		s.Suffix = " " + color.New(color.Bold).Sprint(action)
		// 		s.Start()

		// 		_ = page.Keyboard.Press(input.ControlLeft)
		// 		for _, r := range action.Ctrl {
		// 			if k, ok := keymap[r]; ok {
		// 				_ = page.Keyboard.Type(k)
		// 			}
		// 		}
		// 		_ = page.Keyboard.Release(input.ControlLeft)

		// 		s.Stop()
		// 		fmt.Println(action)
		// 	}
		// }

		// return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
