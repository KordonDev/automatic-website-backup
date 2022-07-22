package web

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/log/v8"
	"golang.org/x/net/html"
)

type Side struct {
	url   *string
	body  *string
	links *[]string
}

func NewSide(url *string, body *string) Side {
	return Side{
		url:   url,
		body:  body,
		links: getLinks(body),
	}
}

func (s *Side) ToString(debug bool) string {
	if debug {
		return fmt.Sprintf("{ url: %s,\nlinks: %s,\nbody: %s}", *s.url, *s.links, *s.body)
	}
	return fmt.Sprintf("{ url: %s,\nlinks: %s}", *s.url, *s.links)
}

func (s *Side) Save() {
	// TODO: Fix pathsü
	err := os.MkdirAll("storage", os.ModePerm)
	if err != nil {
		log.Info(err)
	}
	f, err := os.Create("storage/aktuell.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(*s.body)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Writen to %s", *s.url)

}

func FetchAndParse(base string, url string) *Side {
	resp, err := http.Get(base + url)
	if err != nil {
		log.Error("Could not load page")
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read page")
	}
	body := string(bodyBytes)
	result := NewSide(&url, &body)
	return &result
}

func getLinks(text *string) (data *[]string) {

	tkn := html.NewTokenizer(strings.NewReader(*text))
	var links []string
	var isLink bool

	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			log.Error("Error during tokenizing")
			uniqueLinks := removeDublicates(links)
			return &uniqueLinks
		case tt == html.StartTagToken:
			t := tkn.Token()
			isLink = t.Data == "a"
			sameWebsiteHref := ""
			for _, attr := range t.Attr {
				if attr.Key == "href" && strings.HasPrefix(attr.Val, "/") {
					sameWebsiteHref = attr.Val
				}
			}
			if isLink && sameWebsiteHref != "" {
				links = append(links, sameWebsiteHref)
			}
		case tt == html.TextToken:
			isLink = false
		}
	}
}

func removeDublicates[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
