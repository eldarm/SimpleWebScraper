package scrape

import (
	"container/list"
	"log"
	//"regexp"
	//"strings"
)

func ScrapeSite(urlRoot, localRoot string) {
	base := ParseUrl(urlRoot)
	log.Printf("Processing stie %q, saving to %q", urlRoot, localRoot)
	urls := list.New()
	urls.PushFront(urlRoot)

	// Process urls until done.
	for urls.Len() > 0 {
		// Get next URL to process.
		urlElem := urls.Back()
		urls.Remove(urlElem)
		url := ParseUrl(urlElem.Value.(string))
		if len(url.protocol) == 0 || len(url.host) == 0 {
			// Unnecessary assert.
			log.Printf("Url %q does not have protocol or host: %v", url.url, url)
		}

		// Extract the file.
		b, tp, err := GetHttp(url.Url())
		log.Println(b != nil, tp, err)
		if err != nil {
			log.Fatalf("Cannot get the root url %q with error %v", url.Url(), err)
		}

		if tp == HtmlType { // HTML only processing.
			// Extract links.
			aLinks, imgLinks := ExtractLinks(b)
			log.Printf("Total links: %d Image links: %d", len(aLinks), len(imgLinks))

			// Sanitize links
			aLinksMap := FilterALink(base, url, aLinks)
			log.Println("Local links: ", len(*aLinksMap))
			imgLinksMap := NormalizeImgLink(base, url, imgLinks)
			// for link, ori := range *aLinksMap {
			// 	log.Printf("A' link: %q %q", link, ori)
			// }
			// for link, file := range *imgLinksMap {
			// 	log.Printf("Image link: %q %q", link, file)
			// }
		}
	}
}
