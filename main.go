package main

import (
//"os"
//"fmt"
//"time"
)

const configFilename = "config.toml"
var config Config

func main() {
	config = ReadConfig(configFilename)

	refreshCategories(config.CategoriesUri)
	refreshRegions(config.RegionsUri)

	/*
		for {
			if !areCategoriesFresh(config.MaxAge) {
				fmt.Printf("Categories stale, refreshing!\n")
				refreshCategories(config.CategoriesUri)
				break
			}
			time.Sleep(1 * time.Second)
		}
	*/

	startHttpServer(int(config.Port))
}
