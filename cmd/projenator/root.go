package projenator

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "projenator",
	Short: "projenator - A simple CLI tool used to generate stuff",
}

func Execute() error {
	return rootCmd.Execute()
}
