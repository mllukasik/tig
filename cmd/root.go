package cmd

import (
	"tig/cmd/branch"
	"tig/cmd/push"
	"tig/cmd/version"

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
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(branch.BranchCmd)
	rootCmd.AddCommand(push.PushCmd)
}
