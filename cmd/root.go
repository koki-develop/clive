package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rod/rod/lib/input"
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
		port, err := RandomUnusedPort()
		if err != nil {
			return err
		}

		fmt.Printf("port: %d\n", port)

		ttyd := TTYD(port)
		if err := ttyd.Start(); err != nil {
			return err
		}
		defer ttyd.Process.Kill()

		browser, err := LaunchBrowser()
		if err != nil {
			return err
		}

		page := browser.MustPage(fmt.Sprintf("http://localhost:%d", port))
		_ = page.MustWaitIdle()

		if _, err := page.Eval("() => term.options.fontSize = 32"); err != nil {
			return err
		}

		for _, c := range "echo こんにちは" {
			_ = page.MustElement("textarea").Input(string(c))
			_ = page.MustWaitIdle()
			time.Sleep(100 * time.Millisecond)
		}

		if err := page.Keyboard.Type(input.Enter); err != nil {
			return err
		}

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
