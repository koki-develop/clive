package cmd

import (
	_ "embed"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed clive.yml
var configInitTemplate []byte

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a config file",
	Long:  "Create a config file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(configFilename); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}

			f, err := os.Create(configFilename)
			if err != nil {
				return err
			}
			defer f.Close()

			if _, err := f.Write(configInitTemplate); err != nil {
				return err
			}

			fmt.Printf("created %s\n", configFilename)
			return nil
		}

		return fmt.Errorf("%s already exists", configFilename)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
