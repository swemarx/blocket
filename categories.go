package main

import (
	"fmt"
	"time"
	"strconv"
	"github.com/gocolly/colly"
)

type category struct {
	name string
	id uint64
}

type categories struct {
	list []category
	lastUpdated int64
}

var catList = categories{}

func refreshCategories(uri string) {
	fmt.Println("updateCache(): entering")
	scrapeCategories(uri)

	// DEBUG
	for _, cat := range catList.list {
		fmt.Printf("Name: %s Id: %d\n", cat.name, cat.id)
	}
}

func scrapeCategories(uri string) {
	fmt.Println("scrapeCategories(): entering")

	// Clear out catList.list-slice
	catList.list = nil

	// Scrape 'em
	c := colly.NewCollector()
	c.OnHTML("select.search_category > option", forEachCategory)
	c.Visit(uri)
	catList.lastUpdated = time.Now().Unix()
}

func forEachCategory(e *colly.HTMLElement) {
	// Skip those with data-url attributes defined
	dataurl := e.Attr("data-url")
	if dataurl != "" {
		return
	}

	value := e.Attr("value")
	name  := e.Text
	//fmt.Printf("Found matching option name=%s, value=%s\n", name, value)

	// Some options are just placeholders, dont add those
	if name != "" {
		// Parse category-id
		id, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			fmt.Println("Could not parse category-id")
			return
		}
		catList.list = append(catList.list, category{name: name, id: id})
	}
}

func areCategoriesFresh(maxAge int) bool {
	fmt.Println("areCategoriesFresh(): entering")
	//fi, err := os.Stat(filePath)
	//if err != nil {
	//	return false
	//}
	//timeDiff := time.Now().Unix() - fi.ModTime().Unix()

	timeDiff := time.Now().Unix() - catList.lastUpdated
	if timeDiff > int64(maxAge) {
		return false
	}
	return true
}
