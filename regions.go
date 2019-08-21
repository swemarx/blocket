package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

type region struct {
	Name string
	uri  string
}

type regions struct {
	list        []region
	lastUpdated int64
}

var regList = regions{}

func refreshRegions(uri string) {
	fmt.Println("refreshRegions(): entering")
	scrapeRegions(uri)

	// DEBUG
	fmt.Printf("Discovered regions:\n")
	for _, reg := range regList.list {
		fmt.Printf("name: %s, uri: %s\n", reg.Name, reg.uri)
	}
}

func scrapeRegions(uri string) {
	fmt.Println("scrapeRegions(): entering")

	regList.list = nil

	// Scrape 'em
	c := colly.NewCollector(colly.UserAgent(userAgent))
	c.OnHTML("ul.regionslist > li > a", forEachRegion)
	c.Visit(uri)
	regList.lastUpdated = time.Now().Unix()
}

func forEachRegion(e *colly.HTMLElement) {
	uri := e.Attr("href")
	// Check if the region has subregions, need to scrape that page instead.
	if strings.HasSuffix(uri, ".htm") {
		scrapeSubRegions(uri)
		return
	}
	name := e.Text
	//fmt.Printf("Found potential region name=%s, uri=%s\n", name, uri)
	regList.list = append(regList.list, region{Name: name, uri: uri})
}

func scrapeSubRegions(uri string) {
	fmt.Println("scrapeSubRegions(): entering")

	// Scrape 'em
	c := colly.NewCollector(colly.UserAgent(userAgent))
	c.OnHTML("body > div > a", forEachSubRegion)
	c.Visit(uri)
	regList.lastUpdated = time.Now().Unix()
}

func forEachSubRegion(e *colly.HTMLElement) {
	uri := e.Attr("href")
	name := strings.TrimSpace(e.Text)
	//fmt.Printf("Found potential sub-region name=%s, uri=%s\n", name, uri)
	regList.list = append(regList.list, region{Name: name, uri: uri})
}
