package cmd

import "github.com/spf13/cobra"

func init() {
	for _, cmd := range []*cobra.Command{startCmd, initCmd} {
		cmd.Flags().StringP("config", "c", defaultConfigPath, "config file name")
	}
}
