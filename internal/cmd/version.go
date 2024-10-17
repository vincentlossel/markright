package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0"

// TODO: Add the last commit ID
var lastCommitID = "8c4a11a"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the current version of MarkRight",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Version: %s (%s)", version, lastCommitID))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
