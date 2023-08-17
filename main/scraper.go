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

	scrp.ScrapeSite(os.Args[1], os.Args[2])
}
