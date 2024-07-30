package main

import (
	"github.com/vincentlossel/markright/internal/cmd"
)

func main() {
	initConfig()

	cmd.Execute()
}

func initConfig() {
	setDefaultConfig()
	loadConfig()
}
