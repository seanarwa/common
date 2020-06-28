package main

import (
	"os"

	"seanchang.me/common/config"
)

func main() {

	help_flag, log_config_file, config_file := config.ParseArgs()

	if help_flag {
		config.PrintUsage()
		os.Exit(0)
	}

	config.Set("cmd.log_config_file", log_config_file)
	config.Set("cmd.config_file", config_file)
	config.Init()

	for {
	}
}
