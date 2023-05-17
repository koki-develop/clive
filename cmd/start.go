package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/koki-develop/clive/internal/ui"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start cLive actions",
	Long:  "Start cLive actions.",
	Args:  cobra.MaximumNArgs(0),
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
