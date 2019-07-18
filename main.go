package main

import (
	//"log"
	//"os"
)

const uri = "https://www.blocket.se/stockholm?ca=11"
const filePath = "./data"
const maxAge = 600
const listenPort = 666

func main() {
	// TODO: check if cache exists and is not stale
	if !isCacheFresh(filePath, maxAge) {
		updateCache(uri, filePath)
	}

	// TODO: launch service listening to listenPort
	startServer(listenPort)
}
