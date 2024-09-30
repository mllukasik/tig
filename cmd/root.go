package cmd

import (
	"github.com/spf13/cobra"
	"tig/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:   "tig",
	Short: "Another git tool",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)
}
