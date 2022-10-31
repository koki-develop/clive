package cmd

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "", // TODO
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgname, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		cfg, err := loadConfig(cfgname)
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

		page := browser.
			NoDefaultDevice().
			MustPage(fmt.Sprintf("http://localhost:%d", port)).
			MustWaitIdle()

		_ = page.MustEval("() => term.options.fontSize = 22")
		_ = page.MustEval("term.fit")

		for i, action := range cfg.Actions {
			switch action := action.(type) {
			case *typeAction:
				fmt.Println(action.String())
				for _, c := range action.Type {
					_ = page.MustElement("textarea").Input(string(c))
					_ = page.MustWaitIdle()
					time.Sleep(action.Time)
				}
			case *keyAction:
				fmt.Println(action.String())
				for i := 0; i < action.Count; i++ {
					_ = page.Keyboard.MustType(action.Key)
					time.Sleep(action.Time)
				}
			case *pauseAction:
				next := "quit"
				if i+1 < len(cfg.Actions) {
					next = cfg.Actions[i+1].String()
				}
				fmt.Printf("%s (next: %s)\n", action.String(), next)

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
				fmt.Println(action.String())
				time.Sleep(action.Time)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
