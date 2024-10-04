package branch

import (
	"fmt"
	"github.com/spf13/cobra"
	"tig/git"
	"tig/view/app"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "shows branches",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cmd.Help()
			return
		}
		app.NewApplication().BranchView().Run()
	},
}

var pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "removes all branches",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("specifing branch or remote is not supported yet")
			return
		}
		all, err := cmd.Flags().GetBool("all")
		if all && err != nil {
			git.PruneBranchAll()
			return
		}
		git.PruneBranch()
	},
}

func init() {
	BranchCmd.AddCommand(pruneCmd)
	pruneCmd.Flags().BoolP("all", "a", false, "include current branch")
}
