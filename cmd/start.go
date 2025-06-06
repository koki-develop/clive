package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/clive/internal/ui"
	"github.com/spf13/cobra"
)

// TODO: move to internal/cli
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start cLive actions",
	Long:  "Start cLive actions.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		m := ui.New(flagConfig)
		defer func() { _ = m.Close() }()

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
