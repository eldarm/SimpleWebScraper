package scrape

import (
	"container/list"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	//"regexp"
	//"strings"
)

var (
	maxLinks = flag.Int("max_links", 60000, "maximum number of unique links to rpocess (both A and IMG)"))

func ScrapeSite(urlRoot, localRoot string) {
	base := ParseUrl(urlRoot)
	log.Printf("Processing stie %q, saving to %q", urlRoot, localRoot)

	processed := make(map[string]string, 2000)
	allLinks := make(map[string]*LUrl, 2000) // Contains all encountered links and has their mapping to file names.
	urls := list.New()
	urls.PushFront(urlRoot)

	// Process urls until done.
	for urls.Len() > 0 && urls.Len() < *maxLinks {
		// log.Printf("Links in the queue: %d", urls.Len())

		// Get next URL to process.
		urlElem := urls.Back()
		urls.Remove(urlElem)
		urlString := urlElem.Value.(string)
		processed[urlString] = "Ok"
		url := ParseUrl(urlString)
		allLinks[urlString] = url
		if len(url.protocol) == 0 || len(url.host) == 0 {
			// Unnecessary assert.
			msg := fmt.Sprintf("Url %q does not have protocol or host: %v", url.url, url)
			processed[urlString] = msg
			log.Print(msg)
			continue
		}

		// Get the file.
		b, tp, sfx, err := GetHttp(url.Url())
		//log.Println(b != nil, tp, err)
		if err != nil {
			msg := fmt.Sprintf("Cannot get the root url %q/%q with error %v", url.Url(), urlString, err)
			processed[urlString] = msg
			log.Print(msg)
			continue
		}
		url.suffix = sfx

		// HTML only processing.
		if tp == HtmlType {
			// Extract links.
			aLinks, imgLinks := ExtractLinks(b)
			//log.Printf("Total links: %d Image links: %d", len(aLinks), len(imgLinks))

			// Sanitize links
			aLinksMap := FilterALink(base, url, aLinks)
			imgLinksMap := NormalizeImgLink(base, url, imgLinks)
			//log.Printf("Links: %d Images: %d", len(*aLinksMap), len(*imgLinksMap))
			for link, file := range *imgLinksMap {
				// log.Printf("Image link: %q %q", link, file)
				if _, ok := allLinks[link]; !ok {
					urls.PushFront(link)
					allLinks[link] = file
				}
			}
			for link, file := range *aLinksMap {
				// log.Printf("A' link: %q %q", link, ori)
				if _, ok := allLinks[link]; !ok {
					urls.PushFront(link)
					allLinks[link] = file
				}
			}
			log.Printf("Queue size: %d, new A links: %d, new IMG links: %d",
					   urls.Len(), len(*aLinksMap), len(*imgLinksMap))
		}

		// Save the file
		{
			//err := error(nil)
			url := allLinks[urlString]
			// Need a path
			file := path.Join(localRoot, url.FileName())
			dir, _ := path.Split(file)
			if dir != "" {
				err := os.MkdirAll(dir, 0750)
				//fmt.Printf("  dir: %s err: %v\n", dir, err)
				if err != nil {
					msg := fmt.Sprintf("Failed to create a directory %q", dir)
					processed[urlString] = msg
					log.Print(msg)
					// continue: no, may be we already created it, let the file write fail.
				}
			}
			err = os.WriteFile(file, b, 0660)
			//fmt.Printf(" file: %s err: %v\n", file, err)
			if err != nil {
				msg := fmt.Sprintf("Failed to write a file %q", file)
				processed[urlString] = msg
				log.Print(msg)
				continue
			}
		}
	}
	// Write stats.
	{
	f, err := os.Create(path.Join(localRoot, "all_files_stats.txt"))
	if err != nil {
		log.Printf("Stats file open failure %v.", err)
	}
	defer f.Close()
	for file, e := range processed {
		_, err := f.WriteString(fmt.Sprintf("%s %s\n", file, e))
		if err != nil {
			log.Printf("Stats file write failure %v.", err)
		}
	}
}

	// Write all found links.
	{
		f, err := os.Create(path.Join(localRoot, "all_links_stats.txt"))
		if err != nil {
			log.Printf("Links file open failure %v.", err)
		}
		defer f.Close()
		for urls.Len() >0  {
			url := urls.Front()
			urls.Remove(url) 
			_, err := f.WriteString(fmt.Sprintf("%s\n", url.Value.(string)))
			if err != nil {
				log.Printf("Link file write failure %v.", err)
			}
		}
	}
	
	// for link, url := range allLinks {
	// 		log.Printf("***\nlink: %q\nfile: %q\n", link, url.FileName())
	// }
}
