package scrape

import (
	//"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

)

const (
	DefaultPage = "index.html"
)

var (
	allowArgsFlag = flag.String("allow_args", "page", "process pages with this arg, e.g. q.com/index.php?page=5")
	allowedArgs   = make(map[string]struct{})
)

var (
  //reUrl    = regexp.MustCompile(`(?i:(https?)://([^/#\?]*))?(?:(/?.*)/)?(.*)?$`)
	reUrl    = regexp.MustCompile(`(?i:(https?)://([^/#\?]*))?(?:(/?[^#\?]*)/)?(.*)?$`)
	reSuffix = regexp.MustCompile(`\.[\w\d_-]*$`)
)

type LUrl struct {
	url      string // The original url.
	protocol string // "http" or "hhtps", no colon, no slashes
	host     string // Host, e.g. ibm.com, no slashes.
	path     string // Path, the part after the host and before the file. No start or end slash.
	name     string // The file name, without a leading slash.
	args     string // The part after '?' in the url, if present. Sanitized using allowArgsFlag.
	suffix   string // Suggested suffix, if any
	err      error  // if != nil, only url field is valid.
}

func parseAllowedArgs() {
	if len(*allowArgsFlag) == 0 || len(allowedArgs) > 0 {
		return
	}
	as := strings.Split(*allowArgsFlag, ",")
	for _, a := range as {
		allowedArgs[a] = struct{}{}
	}
}

func ParseUrl(s string) *LUrl {
	parseAllowedArgs() // Not thread safe.
	
	lu := LUrl{
		url: s,
	}

	if st, err := url.PathUnescape(s); err != nil {
		log.Println("Unescape error in path:", s, err)
	} else {
		s = st
	}
	if strings.HasPrefix(s, "//") {
		// Google sites has a bug with "//" but no "http[s]:"
		s = "http:" + s
	}
	s = strings.Replace(s, "index.php/", "", 1) // Some CMS occasionally use index.php as a folder and put it into URL.
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
	if lu.host != strings.Trim(lu.host, "/") {
		log.Panicf("Host contains slashes %s", lu.host)
	}
	lu.path = p[3]
	if lu.path != "" && lu.path[len(lu.path)-1] == '/' {
		log.Panicf(" Path ends with a slash %s", lu.path)
	}

	nm, rs, _ := strings.Cut(p[4], "?") // Empty rs is ok.
	nm, _, _ = strings.Cut(nm, "#") // If not found, that's good.
	lu.name = nm

	// Remove disallowed args.
	args := strings.Split(rs, "&") // If not present, it's "" anyway.
	argsSanitized := make([]string, 0, len(args))
	for _, arg := range args {
		a, _, _ := strings.Cut(arg, "=")
		// log.Println("*** ", arg, a)
		if _, ok := allowedArgs[a]; ok {
			argsSanitized = append(argsSanitized, arg)
			// log.Println("+++ ", argsSanitized)
		}
	}
	if len(argsSanitized) > 0 {
		lu.args = "?" + strings.Join(argsSanitized, "&")
		//log.Println("=== ", lu.args)
	}

	//log.Println(lu.Url())
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
	r := ""
	if lu.protocol != "" {
		r = lu.protocol + "://"
	}
	if lu.host != "" {
		r += strings.Trim(lu.host, "/")
	}
	if lu.path != "" {
		switch r {
		case "":
			r += lu.path
		default:
			r += "/" + strings.Trim(lu.path, "/")
		}
	}
	n := strings.Trim(lu.name, "/")
	if n != "" {
		if r != "" {
			r+= "/"
		}
		r += n
	}
	// return fmt.Sprintf("%s%s/%s%s", r, strings.Trim(lu.path, "/"), lu.name, lu.args)
	return fmt.Sprintf("%s%s", r, lu.args)
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
