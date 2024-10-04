package push

import (
	"fmt"
	"github.com/spf13/cobra"
	"tig/git"
)

var PushCmd = &cobra.Command{
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
