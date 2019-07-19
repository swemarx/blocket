package main

import (
	"fmt"
	"time"
	"strconv"
	"github.com/gocolly/colly"
)

type region struct {
	name string
	id uint64
	uri string
}

type regions struct {
	list []region
	lastUpdated int64
}

var regList = regions{}

func refreshRegions(uri string) {
	fmt.Println("refreshRegions(): entering")
	scrapeRegions(uri)

	// DEBUG
	for _, reg := range regList.list {
		fmt.Printf("Name: %s Id: %d Uri: %s\n", reg.name, reg.id, reg.uri)
	}
}

func scrapeRegions(uri string) {
	fmt.Println("scrapeRegions(): entering")

	regList.list = nil

	// Scrape 'em
	c := colly.NewCollector()
	c.OnHTML("ul.regionslist > li > a", forEachRegion)
	c.Visit(uri)
	regList.lastUpdated = time.Now().Unix()
}

func forEachRegion(e *colly.HTMLElement) {
	name := e.Text
	uri  := e.Attr("href")
	ids  := e.Attr("data-region")
	fmt.Printf("Found matching region name=%s, uri=%s, id=%s\n", name, uri, ids)

	// Parse region-id
	id, err := strconv.ParseUint(ids, 10, 64)
	if err != nil {
		fmt.Println("Could not parse region-id")
		return
	}
	regList.list = append(regList.list, region{name: name, id: id, uri: uri})
}

/*
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
*/
