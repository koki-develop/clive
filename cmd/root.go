package cmd

import (
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var version string

var rootCmd = &cobra.Command{
	Use:   "clive",
	Short: "Automates terminal operations and lets you view them live via a browser",
	Long:  "Automates terminal operations and lets you view them live via a browser.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}

	rootCmd.Version = version

	_ = notifyNewRelease(os.Stderr)
}
