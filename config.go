package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
}

func NewConfig() Config {
	return Config{}
}

func loadConfig() {
	viper.SetConfigName("markright")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("/etc/markright")
	viper.AddConfigPath("$HOME/")
	viper.AddConfigPath("$HOME/.config/markright")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("error config file: %w", err))
	}
}

func setDefaultConfig() {
	viper.SetDefault("templates.source", "~/")
	viper.SetDefault("data.source", "~/.config/markright/")
}
