package cmd

import (
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	version string

	flagConfig string
)

var rootCmd = &cobra.Command{
	Use:   "clive",
	Short: "Automates terminal operations",
	Long:  "Automates terminal operations.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	/*
	 * version
	 */

	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version
	_ = notifyNewRelease(os.Stderr)

	/*
	 * commands
	 */

	rootCmd.AddCommand(
		initCmd,
		startCmd,
		validateCmd,
	)

	/*
	 * flags
	 */

	for _, cmd := range []*cobra.Command{
		startCmd,
		initCmd,
		validateCmd,
	} {
		// --config
		cmd.Flags().StringVarP(&flagConfig, "config", "c", "./clive.yml", "config file name")
	}
}
