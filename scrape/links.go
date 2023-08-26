package scrape

import (
	//"log"
	//"net/url"
	//"regexp"
	"strings"
)

func FilterALink(base, page *LUrl, links []string) *map[string]*LUrl {
	res := make(map[string]*LUrl, len(links))
	for _, link := range links {
		l := ParseUrl(link)
		if l.host == "" || l.protocol == "" { // They are always together empty or not.
			// Local link.
			if l.path != "" && l.path[0] == '/' && !strings.HasPrefix(l.path, base.path) {
				// With absolute path.
				//log.Printf("Sanitized %q, path differs from base %q.", link, base.Url())
				continue
			}
			l := l.Merge(page)
			res[l.Url()] = l
			continue
		}
		// Full link.
		if l.host != base.host {
			//log.Printf("Sanitized %q, host differs from %q.", link, base)
			continue
		}
		if !strings.HasPrefix(l.path, base.path) {
			// With absolute path.
			//log.Printf("Sanitized %q, path differs from base %q.", link, base.Url())
			continue
		}
		res[l.Url()] = l
	}
	return &res
}

func NormalizeImgLink(base, page *LUrl, links []string) *map[string]*LUrl {
	res := make(map[string]*LUrl, len(links))
	for _, link := range links {
		l := ParseUrl(link)
		if l.protocol == "" || l.host == "" {
			if l.path != "" && l.path[0] == '/' {
				l = l.Merge(base)
			} else {
				l = l.Merge(page)
			}
		}
		res[l.Url()] = l
	}
	return &res
}
