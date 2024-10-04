package cmd

import (
	"github.com/mllukasik/tig/cmd/branch"
	"github.com/mllukasik/tig/cmd/push"
	"github.com/spf13/cobra"
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
