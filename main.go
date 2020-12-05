package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/abhishekkr/gol/golhttpclient"
)

const (
	FindUrlBase       = "https://www.imdb.com/find"
	GetParamForMovie  = "ft"
	imdbTitleSelector = ".findList .findResult a"

	HTTPUserAgent = "Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
)

var (
	ExampleTitle = "fightclub"
)

func imdbHttpClient() *golhttpclient.HTTPRequest {
	req := golhttpclient.HTTPRequest{
		GetParams: map[string]string{
			"s":     "tt",
			"ref_":  "fn_ft",
			"ttype": GetParamForMovie,
			"q":     ExampleTitle,
		},
		HTTPHeaders: map[string]string{
			"user-agent": HTTPUserAgent,
		},
	}
	req.Url = FindUrlBase
	golhttpclient.SkipSSLVerify = true
	return &req
}

func main() {
	req := imdbHttpClient()
	response, err := req.Get()
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(imdbTitleSelector).Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Println(title)
	})
}
