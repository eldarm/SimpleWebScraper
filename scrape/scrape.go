package scrape

import (
	"container/list"
	"log"
	"regexp"
	"strings"
)

var (
	reFileName = regexp.MustCompile(`/[^/]+$`)
	reProtocol = regexp.MustCompile(`^(?i)http[s]?://`)
)

func ScrapeSite(urlRoot, localRoot string) {
	log.Printf("Processing stie %q, saving to %q", urlRoot, localRoot)
	urls := list.New()
	urls.PushFront(urlRoot);

	// Process urls until done.
	for urls.Len() > 0 {
		// Get next URL to process.
		urlElem := urls.Back()
		urls.Remove(urlElem)
		url := urlElem.Value.(string)
		if len(reProtocol.FindAllString(url, -1)) == 0 {
			// Unnecessary assert.
			log.Printf("Url %q does not have protocol.", url)
		}

		// Extract the file.
		b, tp, err := GetHttp(url)
		log.Println(b != nil, tp, err)
		if err != nil {
			log.Fatalf("Cannot get the root url %q with error %v", url, err)
		}
	
		if tp == HtmlType { // HTML only processing.
			// Extract links.
			aLinks, imgLinks := ExtractLinks(b)
			log.Printf("Total links: %d Image links: %d", len(aLinks), len(imgLinks))
	
			// Normalize links
			aLinks = Filter(url, aLinks)
			log.Println("Local links: ", len(aLinks))
			for i, link := range aLinks {
				log.Printf("Link: %d %q", i, link)
			}
			for i, link := range imgLinks {
				log.Printf("Link: %d %q", i, reFileName.FindAllString(strings.Trim(link,"/"), 1))
			}
		}

	}
}