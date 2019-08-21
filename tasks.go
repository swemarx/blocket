package main

import (
//"fmt"
//"time"
//"strconv"
//"github.com/gocolly/colly"
)

type task struct {
	name        string
	region      string
	category    int64
	query       string
	notify      string
	lastUpdated int64
}

type tasks struct {
	list        []task
	lastUpdated int64
}

var taskList = tasks{}
