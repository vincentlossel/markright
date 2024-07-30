package cmd

import (
	"fmt"

	"github.com/vincentlossel/markright/internal/services"

	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "templates",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		allTemplates, err := services.ListTemplates()
		if err != nil {
			// TODO: Handle errors
		}

		fmt.Println("allTemplates:", allTemplates)
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
