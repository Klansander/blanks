package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://github.com/OP-Engineering/link-preview-js св cd")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		rel, _ := s.Attr("rel")
		if rel == "icon" || rel == "shortcut icon" {
			href, _ := s.Attr("href")
			fmt.Println("Favicon URL:", href)
		}
	})
}
