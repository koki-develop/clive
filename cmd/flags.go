package cmd

import "github.com/spf13/cobra"

var (
	configFilename string
)

func init() {
	for _, cmd := range []*cobra.Command{startCmd, initCmd} {
		cmd.Flags().StringVarP(&configFilename, "config", "c", "./clive.yml", "config file name")
	}
}
