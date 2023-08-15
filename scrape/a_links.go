package scrape

import (
	"log"
	"strings"
	"regexp"
	"net/url"
)

var (
	rePrefix = regexp.MustCompile(`^http[s]?://`)
	//reAnchor = regexp.MustCompile(`(?:#|\%23).*`)
	reAnchor = regexp.MustCompile(`#.*`)
	reParams = regexp.MustCompile(`\?.*`)
)

func clean(link string) string {
	if len(link) == 0 || link[0] == '?' {
		return link
	}
	return strings.Trim(reParams.ReplaceAllString(link, ""), "/")
}

func cleanExternalLink(base, link string) string  {
	link, err := url.PathUnescape(link)
	if err != nil {
		log.Printf("Malformed URL %q, error: %s", link, err)
		return ""
	}
	link = strings.Trim(reAnchor.ReplaceAllString(strings.ToLower(link), ""), "/")
	if len(link) == 0 {
		return ""
	}
	if r := rePrefix.FindString(link); len(r) == 0 && len(link) > 0 {
		// Local link without http[s]:// prefix.
		link = clean(link)
		if len(link) > 0 {
			//log.Println("Added a link with no http:// prefix: ", link)
			return link
		}
		return ""
	}
	link = rePrefix.ReplaceAllString(link, "")
	//log.Println("Checking link: ", link)
	if strings.HasPrefix(link, base) {
		// Local link, the starting path matches.
		l := clean(link[len(base):])
		if len(l) > 0 {
			// No empty or inside the page links.
			//log.Println("Added a local link: ", l)
			return l
		}
	}
	//log.Println("Sanitized ", link)
	return ""
}

func Filter(url string, links []string) []string {
	res := make(map[string]bool, len(links))
	//urlClean := strings.Trim(rePrefix.ReplaceAllString(strings.ToLower(url), ""), "/")
	urlClean := rePrefix.ReplaceAllString(strings.ToLower(url), "")
	log.Println("Base url is ", urlClean)
	for _, link := range links {
		l := cleanExternalLink(urlClean, link)
		if len(l) == 0 {
			continue
		}
		res[l] = true
	}
	r := make([]string, 0, len(res))
	for key := range res {
		r = append(r, key)
	}
	return r
} 