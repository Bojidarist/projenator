package projenator

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("1.0.0")
		return nil
	},
}
