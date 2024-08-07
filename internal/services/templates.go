package services

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/viper"
)

func ListTemplates() ([]fs.DirEntry, error) {
	templateDir := viper.Get("templates.source")

	files, err := os.ReadDir(templateDir.(string))
	if err != nil {
		fmt.Println(err)
	}

	return files, nil
}

func GetTemplate(fileName string) (fs.DirEntry, error) {
	allTemplates, err := ListTemplates()
	if err != nil {
		return nil, err
	}

	// TODO: Find a better way to retrieve a template
	for _, file := range allTemplates {
		if file.Name() == fileName {
			return file, nil
		}
	}

	return nil, nil
}

// TODO: Find any placeholder value (e.g. `<% DATE %>`) in the template, and replace it with the corresponding data
// Some pieces of data are contained in a JSON data file, others are dynamic values calculated in time (e.g. DATE, TODAY...)
func ParseTemplate(content []byte) ([]byte, error) {
	return content, nil
}
