package scrape

import (
	//"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

const (
	DefaultPage = "index.html"
)

var (
	// reFileName = regexp.MustCompile(`/[^/]+$`)
	// reProtocol = regexp.MustCompile(`^(?i)http[s]?://`)
	//reUrl = regexp.MustCompile(`(?i:(https?)://([^/]*)/)?(?:(.*)/)?([^#\?]]*)?(?:#[^?]*)?(\?.*)?$`)
	//reUrl = regexp.MustCompile(`(?i:(https?)://([^/]*))?(?:/(.*)/)?(.*)?$`)
	reUrl    = regexp.MustCompile(`(?i:(https?)://([^/]*))?(?:(/?.*)/)?(.*)?$`)
	reSuffix = regexp.MustCompile(`\.[\w\d_-]*$`)
)

type LUrl struct {
	url      string // The original url.
	protocol string // "http" or "hhtps", no colon, no slashes
	host     string // Host, e.g. ibm.com, no slashes.
	path     string // Path, the part after the host and before the file. No start or end slash.
	name     string // The file name, without a leading slash.
	args     string // The part after '?' in the url, if present.
	suffix   string // Suggested suffix, if any
	err      error  // if != nil, only url field is valid.
}

func ParseUrl(s string) *LUrl {
	lu := LUrl{
		url: s,
	}
	if strings.HasPrefix(s, "//") {
		// Google has a bug with "//" but no "http[s]:"
		s = "http:" + s
	}
	pp := reUrl.FindAllStringSubmatch(s, -1)
	if len(pp) < 1 {
		lu.err = fmt.Errorf("url %q has invalid format, zero submatches", s)
		return &lu
	}
	p := pp[0]
	// log.Printf("%q %v", s, p)
	if len(p) != 5 {
		lu.err = fmt.Errorf("url %q has invalid format, parse result: %v", s, p)
		return &lu
	}
	lu.protocol = p[1]
	lu.host = p[2]
	lu.path = p[3]
	// if lu.host != "" { // || s[0] == '/' {
	// 	// Not a relative path:
	// 	lu.path = "/" + lu.path
	// }
	// Cut away anchor, split file name and args.
	nm, rs, r := strings.Cut(p[4], "#")
	if !r {
		nm, rs, _ = strings.Cut(p[4], "?")
	} else {
		_, rs, _ = strings.Cut(rs, "?")
	}
	lu.name = nm
	lu.args = rs // If not present, it's "" anyway.
	return &lu
}

func (lu *LUrl) Copy() *LUrl {
	r := *lu
	return &r
}

func (lu *LUrl) Merge(o *LUrl) *LUrl {
	r := lu.Copy()
	if r.protocol == "" {
		r.protocol = o.protocol
	}
	if r.host == "" {
		r.host = o.host
	}
	if r.path != "" && r.path[0] != '/' && o.path != "" {
		r.path = fmt.Sprintf("%s/%s", o.path, r.path)
	}
	return r
}

func (lu *LUrl) Url() string {
	return fmt.Sprintf("%s://%s/%s/%s", lu.protocol, strings.Trim(lu.host, "/"), strings.Trim(lu.path, "/"), lu.name)
}

func (lu LUrl) String() string {
	return fmt.Sprintf("%s Err: %v",
		strings.Join([]string{lu.url, lu.protocol, lu.host, lu.path, lu.name, lu.args}, " -- "), lu.err)
}

// Convert URl to a relative file path with file name to save.
func (lu LUrl) FileName() string {
	if lu.host == "" {
		log.Fatalf("No host field for %q url aka %q.", lu.url, lu.Url())
	}
	name := lu.name
	if name == "" {
		name = DefaultPage
	}
	sfx := ""
	if len(reSuffix.Find([]byte(name))) == 0 {
		sfx = lu.suffix
	}
	if lu.path == "" {
		return fmt.Sprintf("%s/%s%s", lu.host, name, sfx)
	}
	return fmt.Sprintf("%s/%s/%s%s", strings.Trim(lu.host, "/"), strings.Trim(lu.path, "/"), name, sfx)
}
