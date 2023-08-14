package main

import (
	scrp "SimpleWebScraper/scrape"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello!", os.Args)
	if len(os.Args) < 3 {
		fmt.Println(`
	Usage: scraper <url> <local path>
	url: website or website path, e.g. http://sample.org/mypath/.
	local path: local disk path to save the website.
	`)
		log.Fatalf("The number of command line args is less than two: %v", os.Args)
	}
	url, _ := os.Args[1], os.Args[2]
	b, t, err := scrp.GetHttp(url)
	log.Println(b != nil, t, err)
	if err != nil {
		log.Fatalf("Cannot get the root url %q with error %v", url, err)
	}
	aLinks, imgLinks := scrp.ExtractLinks(b)
	log.Printf("Total links: %d Image links: %d", len(aLinks), len(imgLinks))
	aLinks = scrp.Filter(url, aLinks)
	log.Println("Local links: ", len(aLinks))
	for i, link := range aLinks {
		log.Printf("Link: %d %q", i, link)
	}
	for i, link := range imgLinks {
		log.Printf("Link: %d %q", i, link)
	}
}

