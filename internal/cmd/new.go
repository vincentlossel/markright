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

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("require at least one argument")
		}

		// TODO: Check whether the template passed is valid or not

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Get the template from args
		template, err := services.GetTemplate("Daily Review.md")
		if err != nil {
			// TODO: Handle errors
			fmt.Println(err)
		}

		content, err := os.ReadFile(fmt.Sprintf("%s/%s", viper.Get("templates.source").(string), template.Name()))
		if err != nil {
			fmt.Println(err)
		}

		// TODO: Parse the template, and return the updated content
		parsedContent, err := services.ParseTemplate(content)
		if err != nil {
			fmt.Println(err)
		}

		// TODO: Prefer the user input, if a filename has been provided
		newFileName := services.GenerateFileName(template.Name())

		target := viper.Get("templates.target").(string)
		if err = os.WriteFile(fmt.Sprintf("%s/%s.md", target, newFileName), parsedContent, fs.ModePerm); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
