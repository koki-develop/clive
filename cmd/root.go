package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clive",
	Short: "", // TODO
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig("clive.yml")
		if err != nil {
			return err
		}

		port, err := randomUnusedPort()
		if err != nil {
			return err
		}

		ttyd := ttyd(port)
		if err := ttyd.Start(); err != nil {
			return err
		}
		defer ttyd.Process.Kill()

		browser, err := launchBrowser()
		if err != nil {
			return err
		}

		page := browser.MustPage(fmt.Sprintf("http://localhost:%d", port))
		_ = page.MustWaitIdle()

		if _, err := page.Eval("() => term.options.fontSize = 22"); err != nil {
			return err
		}

		for _, action := range cfg.Actions {
			switch action := action.(type) {
			case *typeAction:
				for _, c := range action.Type {
					_ = page.MustElement("textarea").Input(string(c))
					_ = page.MustWaitIdle()
					time.Sleep(action.Time)
				}
			case *keyAction:
				for i := 0; i < action.Count; i++ {
					_ = page.Keyboard.MustType(action.Key)
					time.Sleep(action.Time)
				}
			case *pauseAction:
				for {
					_, key, err := keyboard.GetSingleKey()
					if err != nil {
						return err
					}
					if key == keyboard.KeyEnter {
						break
					}
				}
			case *sleepAction:
				time.Sleep(action.Time)
			}
		}

		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
