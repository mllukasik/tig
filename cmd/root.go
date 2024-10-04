package cmd

import (
	"github.com/spf13/cobra"
	"tig/cmd/branch"
	"tig/cmd/push"
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
	rootCmd.AddCommand(branch.BranchCmd)
	rootCmd.AddCommand(push.PushCmd)
}
