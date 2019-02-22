package main

import (
	"GoTutorial/advancedGo/handlers"
	"fmt"
)

type location struct {
	loc string `xml:"loc"`
}
type siteMapIndex struct {
	locations []location `xml:"sitemap"`
}

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.ListenAndServe(":8000", nil)
	fmt.Println("hai")

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	var s siteMapIndex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.locations)
	fmt.Println(handlers.Index)
}
