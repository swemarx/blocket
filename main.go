package main

import (
	"fmt"
	"os"
	/*
		"time"
	*/)

const configFilename = "config.toml"

var config Config

func main() {
	config = readConfig(configFilename)
	db, err := openSqlite(config.Sqlite)
	if err != nil {
		fmt.Printf("error: could not open sqlite-db %s\n", config.Sqlite)
		os.Exit(1)
	}

	// TODO: actually use the db
	closeSqlite(db)

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
