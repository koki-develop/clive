package cmd

import "github.com/spf13/cobra"

var (
	flagConfig string
)

func init() {
	for _, cmd := range []*cobra.Command{startCmd, initCmd, validateCmd} {
		cmd.Flags().StringVarP(&flagConfig, "config", "c", "./clive.yml", "config file name")
	}
}
