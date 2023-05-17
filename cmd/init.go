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
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		exists, err := util.Exists(configFilename)
		if err != nil {
			return err
		}
		if exists {
			return fmt.Errorf("%s already exists", configFilename)
		}

		f, err := util.CreateFile(configFilename)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write(configInitTemplate); err != nil {
			return err
		}

		_, _ = fmt.Printf("Created %s\n", configFilename)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
