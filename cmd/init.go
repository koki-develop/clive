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
		if _, err := os.Stat(defaultConfigPath); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}

			f, err := os.Create(defaultConfigPath)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := f.Write([]byte(configInitTemplate)); err != nil {
				return err
			}

			fmt.Printf("created %s\n", defaultConfigPath)
			return nil
		}

		return fmt.Errorf("%s already exists", defaultConfigPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
