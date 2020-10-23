package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func startHttpServer(listenPort int) {
	// Setup file-serving for "assets/" directory
	fs := http.FileServer(http.Dir("html/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	// Setup templating for main page
	indexTmpl := template.Must(template.ParseFiles("html/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			SiteName   string
			PageTitle  string
			Categories []category
			Regions    []region
		}{
			SiteName:   config.Sitename,
			PageTitle:  "Blocko Loco",
			Categories: catList.list,
			Regions:    regList.list,
		}
		err := indexTmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
	})

	fmt.Printf("startHttpServer(): listening to port %d\n", listenPort)
	http.ListenAndServe(":"+strconv.Itoa(listenPort), nil)
}
