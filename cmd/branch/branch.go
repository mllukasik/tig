package branch

import (
	"fmt"
	"tig/git"
	"tig/view/app"

	"github.com/spf13/cobra"
)

var BranchCmd = &cobra.Command{
	Use:   "branch",
	Short: "shows branches",
	Run: func(cmd *cobra.Command, args []string) {
		app.NewApplication().BranchView().Run()
	},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push current branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println("specifing branch or remote is not supported yet")
			return
		}
		git.PushCurrent("origin")
	},
}

func init() {
	BranchCmd.AddCommand(pushCmd)
}
