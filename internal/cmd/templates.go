package cmd

import (
	"fmt"
	"strings"

	"github.com/vincentlossel/markright/internal/services"

	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "templates",
	Short: "Get the list of available templates",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		allTemplates, err := services.ListTemplates()
		if err != nil {
			// TODO: Handle errors
			fmt.Println(err)
		}

		if len(allTemplates) < 1 {
			fmt.Println("No template available")
			return
		}

		// TODO: Filter valid templates
		// TODO: Show default template

		fmt.Println("Available templates:")
		for _, template := range allTemplates {
			fileName := strings.ReplaceAll(template.Name(), ".md", "")
			actionableFileName := services.GetActionableFileName(fileName)

			fmt.Println("*", actionableFileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
}
