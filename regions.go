package main

import (
	"fmt"
	"time"
	"strings"
	"github.com/gocolly/colly"
)

type region struct {
	name string
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
		fmt.Printf("name: %s, uri: %s\n", reg.name, reg.uri)
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
	uri  := e.Attr("href")
	// Check if the region has subregions, need to scrape that page instead.
	if strings.HasSuffix(uri, ".htm") {
		scrapeSubRegions(uri)
		return
	}
	name := e.Text
	fmt.Printf("Found matching region name=%s, uri=%s\n", name, uri)
	regList.list = append(regList.list, region{name: name, uri: uri})
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

func scrapeSubRegions(uri string) {
	fmt.Println("scrapeSubRegions(): entering")

	// Scrape 'em
	c := colly.NewCollector()
	c.OnHTML("body > div > a", forEachSubRegion)
	c.Visit(uri)
	regList.lastUpdated = time.Now().Unix()
}

func forEachSubRegion(e *colly.HTMLElement) {
	uri  := e.Attr("href")
	name := strings.TrimSpace(e.Text)
	fmt.Printf("Found matching region name=%s, uri=%s\n", name, uri)
	regList.list = append(regList.list, region{name: name, uri: uri})
}

