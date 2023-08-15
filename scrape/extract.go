package scrape

import (
	"bytes"
	///"fmt"
	//"io"
	"log"
	"golang.org/x/net/html"
)

// Returns all A-links and IMG-links from HTML body.
func ExtractLinks(body []byte) ([]string, []string) {
	aLinks := make([]string, 0, 200)
	imgLinks := make([]string, 0, 200)
	p := html.NewTokenizer(bytes.NewReader(body))
	for {
		token := p.Next()
		// log.Println("Ended with error token: ", token)
		switch token {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			log.Println("Token: ", token)
			return aLinks, imgLinks
		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken:	
			//st := token
			token := p.Token()
			//log.Printf("<%s>", token.Data)
			switch token.Data{
			case "a":
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						//log.Println("*** ", st, attr.Val)
						aLinks = append(aLinks, attr.Val)
					}
				}
			case "img":
				//log.Printf("<%s>", token.Data)
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						//log.Println("*** ", attr.Val)
						imgLinks = append(imgLinks, attr.Val)
					}
				}
			}
		}
	}
}
