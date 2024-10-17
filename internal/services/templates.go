package services

import (
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"time"

	"github.com/spf13/viper"
)

func ListTemplates() ([]fs.DirEntry, error) {
	templateDir := viper.Get("templates.source")

	files, err := os.ReadDir(templateDir.(string))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return files, nil
}

func LoadTemplates() (map[string]string, error) {
	allTemplates, err := ListTemplates()
	if err != nil {
		return nil, err
	}

	availableTemplates := make(map[string]string)
	for _, template := range allTemplates {
		actionableFileName := GetActionableFileName(template.Name())
		availableTemplates[actionableFileName] = template.Name()
	}

	return availableTemplates, nil
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

func GetTemplateFromActionableName(actionableName string) (fs.DirEntry, error) {
	availableTemplates, err := LoadTemplates()
	if err != nil {
		return nil, err
	}

	// TODO: Check that `actionableName` is in `availableTemplates` or it will SEGFAULT
	template, err := GetTemplate(availableTemplates[actionableName])
	if err != nil {
		return nil, err
	}

	return template, nil
}

// TODO: Find any placeholder value (e.g. `<% DATE %>`) in the template, and replace it with the corresponding data
// Some pieces of data are contained in a JSON data file, others are dynamic values calculated in time (e.g. DATE, TODAY...)
func ParseTemplate(content []byte) ([]byte, error) {
	text := string(content)
	re := regexp.MustCompile(`<%([^%]+)%>`)
	result := re.ReplaceAllStringFunc(text, func(match string) string {
		fmt.Println("MATCH", match)
		switch match {
		case "<% DATE %>":
			now := time.Now()
			today := fmt.Sprintf("%d-%d-%d", now.Year(), now.Month(), now.Day())
			return today
		default:
			return match
		}

	})

	parsedContent := []byte(result)

	return parsedContent, nil
}
