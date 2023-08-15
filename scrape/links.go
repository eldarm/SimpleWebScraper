package scrape

import (
	"log"
	//"net/url"
	"regexp"
	"strings"
)

var (
	reFileName = regexp.MustCompile(`/[^/]+$`)
	//reProtocol = regexp.MustCompile(`^(?i)http[s]?://`)
)

func FilterALink(base, page *LUrl, links []string) *map[string]string {
	res := make(map[string]string, len(links))
	for _, link := range links {
		l := ParseUrl(link)
		if l.host == "" || l.protocol == "" { // They are always together empty or not.
			// Local link.
			if l.path[0] == '/' && !strings.HasPrefix(l.path, base.path) {
				// With absolute path.
				log.Printf("Sanitized %q, path differs from base %q.", link, base.Url())
				continue
			}
			r := l.Merge(page)
			res[r.Url()] = link
			continue
		}
		// Full link.
		if l.host != base.host {
			log.Printf("Sanitized %q, host differs from %q.", link, base)
			continue
		}
		if !strings.HasPrefix(l.path, base.path) {
			// With absolute path.
			log.Printf("Sanitized %q, path differs from base %q.", link, base.Url())
			continue
		}
	}
	return &res
}

func NormalizeImgLink(base, page *LUrl, links []string) *map[string]string {
	res := make(map[string]string, len(links))
	for _, link := range links {
		l := ParseUrl(link)
		if l.protocol == "" || l.host == "" {
			if l.path[0] == '/' {
				l = l.Merge(base)
			} else {
				l = l.Merge(page)
			}
		}
		fName := reFileName.FindAllString(strings.Trim(link, "/"), 1)[0]
		res[l.Url()] = fName
	}
	return &res
}