package scrape

import (
	///"fmt"
	"io"
	"log"
	"golang.org/x/net/html"
)

// Returns A-links and IMG-links.
func ExtractLinks(body io.ReadCloser) ([]string, []string) {
	aLinks := make([]string, 0, 200)
	imgLinks := make([]string, 0, 200)
	p := html.NewTokenizer(body)
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
				log.Printf("<%s>", token.Data)
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
