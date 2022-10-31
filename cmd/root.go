package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/koki-develop/clive/pkg/helper"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clive",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		port, err := helper.RandomUnusedPort()
		if err != nil {
			return err
		}

		fmt.Printf("port: %d\n", port)

		ttyd := helper.TTYD(port)
		if err := ttyd.Start(); err != nil {
			return err
		}
		defer ttyd.Process.Kill()

		time.Sleep(10 * time.Second)

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
