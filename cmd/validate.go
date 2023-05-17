package cmd

import (
	"fmt"

	"github.com/koki-develop/clive/internal/config"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a config file",
	Long:  "Validate a config file.",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := config.Load(configFilename); err != nil {
			return err
		}

		_, _ = fmt.Printf("Config file %s is valid!\n", configFilename)
		return nil
	},
}
