package main

import (
	//"os"
	//"fmt"
	//"time"
)

const categoriesUri = "https://www.blocket.se/stockholm?ca=11"
const regionsUri = "https://www.blocket.se/"
const maxAge = 10	// seconds
const listenPort = 666

func main() {
	refreshCategories(categoriesUri)
	refreshRegions(regionsUri)

	/*
	for {
		if !areCategoriesFresh(maxAge) {
			fmt.Printf("Categories stale, refreshing!\n")
			refreshCategories(categoriesUri)
			break
		}
		time.Sleep(1 * time.Second)
	}
	*/

	startServer(listenPort)
}
