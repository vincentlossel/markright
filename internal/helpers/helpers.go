package helpers

import "github.com/spf13/viper"

func GetTemplatesPath() string {
	templatesDir := viper.Get("templates.source")
	templatesPath, ok := templatesDir.(string)
	if !ok {
		return ""
	}

	return templatesPath
}
