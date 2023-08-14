package scrape

import (
	//"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	ErrType = iota
	HtmlType
	OtherType
)

func GetHttp(url string) (io.ReadCloser, int, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Cannot download %s with error %s", url, err)
		// TODO: add to the list of errors or?
		return nil, ErrType, err
	}
	//defer resp.Body.Close()
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Printf("Cannot read body of %s with error %s", url, err)
	// 	// TODO: add to the list of errors or?
	// 	return "", ErrType, err
	// }
	contentType := OtherType
	log.Println("Content-type: ", resp.Header.Get("Content-type"))
	if ct := resp.Header.Get("Content-type"); strings.HasPrefix(ct, "text/html") {
		contentType = HtmlType
	}
	return resp.Body, contentType, nil
	
}
