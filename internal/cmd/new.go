package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/vincentlossel/markright/internal/services"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var templateName string
var addTimestamp bool

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new file",
	Long:  "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require at least one argument")
		}

		// TODO: Check whether the template passed is valid or not

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Return an error if the file already exists (especially if there is no timestamp)

		template, err := services.GetTemplateFromActionableName(templateName)
		if err != nil {
			// TODO: Handle errors
			fmt.Println("ERROR:", err)
		}

		sourceDir, ok := viper.Get("templates.source").(string)
		if !ok {
			fmt.Println("ERROR: Could not find sourceDir in `templates.source`")
			return
		}

		content, err := os.ReadFile(fmt.Sprintf("%s/%s", sourceDir, template.Name()))
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		// TODO: Parse the template, and return the updated content
		parsedContent, err := services.ParseTemplate(content)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		// Prefer the user input, if a filename has been provided
		var newFileName string
		if args[0] != "" {
			newFileName = services.GenerateFileName(args[0], addTimestamp)
		} else {
			newFileName = services.GenerateFileName(template.Name(), addTimestamp)
		}

		// TODO: Get the target from the frontmatter of the template, if provided
		target := viper.Get("templates.target").(string)
		if err = os.WriteFile(fmt.Sprintf("%s/%s.md", target, newFileName), parsedContent, fs.ModePerm); err != nil {
			fmt.Println("ERROR:", err)
			return
		}
	},
}

func init() {
	// TODO: Get the default parameters from the configuration file
	newCmd.Flags().StringVarP(&templateName, "template", "t", "DailyReview", "Template Name")
	newCmd.Flags().BoolVarP(&addTimestamp, "timestamp", "d", false, "Timestamp")

	rootCmd.AddCommand(newCmd)
}
