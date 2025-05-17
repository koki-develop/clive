package cmd

import (
	"os"

	"github.com/koki-develop/clive/internal/cli"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a config file",
	Long:  "Create a config file.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		c := cli.New(&cli.Config{
			Stdout: os.Stdout,
		})

		return c.Init(&cli.InitParams{
			Config: flagConfig,
		})
	},
}
