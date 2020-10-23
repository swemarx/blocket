package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Sitename      string
	Useragent     string
	CategoriesUri string
	RegionsUri    string
	MaxAge        uint
	Port          uint
	Debug         bool
	Sqlite        string
}

// Reads info from config file
func readConfig(configFilename string) Config {
	_, err := os.Stat(configFilename)
	if err != nil {
		fmt.Println("Config file is missing: ", configFilename)
	}

	if _, err := toml.DecodeFile(configFilename, &config); err != nil {
		fmt.Println(err)
	}

	return config
}
