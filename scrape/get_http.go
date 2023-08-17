package scrape

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	ErrType = iota
	HtmlType
	ImageType
	OtherType
)

func GetHttp(url string) ([]byte, int, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Cannot download %s with error %s", url, err)
		// TODO: add to the list of errors or?
		return nil, ErrType, "", err
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	if code >= 300 {
		e := fmt.Errorf("cannot download %s with http status code %d error %s", url, code, resp.Status)
		log.Print(e)
		return nil, ErrType, "", e
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Cannot read body of %s with error %s", url, err)
		// TODO: add to the list of errors or?
		return nil, ErrType, "", err
	}

	if bytes.Contains(body[:3], []byte("404")) {
		log.Printf("Cannot read body of %s with 404 error %s", url, string(body[30]))
		// TODO: add to the list of errors or?
		return nil, ErrType, "", err
	}
	contentType := OtherType
	suggestedSuffix := ""
	//log.Println("Content-type: ", resp.Header.Get("Content-type"))
	ct := resp.Header.Get("Content-type")
	switch {
	case strings.HasPrefix(ct, "text/html"):
		contentType = HtmlType
		suggestedSuffix = ".html"
	case strings.HasPrefix(ct, "image/"):
		contentType = ImageType
		v := strings.Split(ct, "/")
		if len(v) > 1 {
			suggestedSuffix = "." + v[1]
		}
	}
	return body, contentType, suggestedSuffix, nil
}
