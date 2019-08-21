package main

import (
//"os"
//"fmt"
//"time"
)

const siteName = "Blocko Loco"
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"
const categoriesUri = "https://www.blocket.se/stockholm?ca=11"
const regionsUri = "https://www.blocket.se/"
const maxAge = 10 // seconds
const listenPort = 8081

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

	startHttpServer(listenPort)
}
