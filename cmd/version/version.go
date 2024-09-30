package version

import (
	"fmt"
	"tig/build"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", build.Version)
	},
}
