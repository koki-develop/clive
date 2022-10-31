package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const configInitTemplate = `actions:
  - pause
  - type: echo 'Welcome to clive!'
  - key: enter
  - pause
`

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "", // TODO
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			return err
		}

		if _, err := os.Stat(config); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}

			f, err := os.Create(config)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := f.Write([]byte(configInitTemplate)); err != nil {
				return err
			}

			fmt.Printf("created %s\n", config)
			return nil
		}

		return fmt.Errorf("%s already exists", config)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
