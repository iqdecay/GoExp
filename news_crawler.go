package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Loc string `xml:"loc"`
}
type Sitemapindex struct {
	Locations []Location `xml:"sitemap"`
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	var s Sitemapindex
	xml.Unmarshal(bytes, &s)
	fmt.Println(s.Locations)
}
