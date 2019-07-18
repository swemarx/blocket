package main

import (
	"log"
	"os"
	"time"
	"strconv"
	"github.com/gocolly/colly"
)

type category struct {
	name string
	id uint64
}

var categories = []category{}

func updateCache(uri string, filePath string) {
	log.Printf("updateCache(): entering")
	scrapeSite(uri)

	// DEBUG
	for _, cat := range categories {
		log.Printf("Name: %s Id: %d\n", cat.name, cat.id)
	}
}

func scrapeSite(uri string) {
	log.Printf("scrapeSite(): entering")

	// Clear out categories-slice
	categories = nil

	// Scrape it
	c := colly.NewCollector()
	c.OnHTML("select.search_category > option", forEachCategory)
	c.Visit(uri)
}

func forEachCategory(e *colly.HTMLElement) {
	// Skip those with data-url attributes defined
	dataurl := e.Attr("data-url")
	if dataurl != "" {
		return
	}

	value := e.Attr("value")
	name  := e.Text
	//log.Printf("Found matching option name=%s, value=%s\n", name, value)

	// Some options are just placeholders, dont add those
	if name != "" {
		// Parse category-id
		id, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			log.Printf("Could not parse category-id")
			return
		}
		categories = append(categories, category{name: name, id: id})
	}
}

// Returns true if file exists and mtime < maxAge, otherwise false
func isCacheFresh(filePath string, maxAge int) bool {
	log.Printf("isCacheFresh(): entering")
	fi, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	timeDiff := time.Now().Unix() - fi.ModTime().Unix()
	if timeDiff > int64(maxAge) {
		return false
	}
	return true
}
