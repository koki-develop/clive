package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "clive",
	Version: "v0.1.0",
	Short:   "Automates terminal operations and lets you view them live via a browser",
	Long:    "Automates terminal operations and lets you view them live via a browser.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
