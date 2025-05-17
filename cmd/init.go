package cmd

import (
	_ "embed"
	"fmt"

	"github.com/koki-develop/clive/internal/util"
	"github.com/spf13/cobra"
)

//go:embed clive.yml
var configInitTemplate []byte

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create a config file",
	Long:  "Create a config file.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		exists, err := util.FileExists(flagConfig)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("%s already exists", flagConfig)
		}

		f, err := util.CreateFile(flagConfig)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write(configInitTemplate); err != nil {
			return err
		}

		_, _ = fmt.Printf("Created %s\n", flagConfig)
		return nil
	},
}
