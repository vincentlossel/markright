package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("Version: %s", version))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
