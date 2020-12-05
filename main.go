package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/abhishekkr/gol/golhttpclient"
)

const (
	ImdbUrlBase                 = "https://www.imdb.com"
	FindUrlBase                 = "https://www.imdb.com/find"
	GetParamForMovie            = "ft"
	imdbTitleSelector           = ".findList .findResult .result_text a"
	creditSummaryLabelsSelector = ".credit_summary_item h4"

	HTTPUserAgent = "Mozilla/5.0 (X11; Fedora; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
)

var (
	ExampleTitle = "fightclub"
)

type Movie struct {
	Name      string
	Link      string
	Directors []Artist
}

type Artist struct {
	Name string
	Link string
}

func imdbHttpClient() *golhttpclient.HTTPRequest {
	req := golhttpclient.HTTPRequest{
		GetParams: map[string]string{
			"s":    "tt",
			"ref_": "fn_ft",
		},
		HTTPHeaders: map[string]string{
			"user-agent": HTTPUserAgent,
		},
	}
	golhttpclient.SkipSSLVerify = true
	return &req
}

func findMovie(query string) []Movie {
	req := imdbHttpClient()
	req.Url = FindUrlBase
	req.GetParams["ttype"] = GetParamForMovie
	req.GetParams["q"] = query
	response, err := req.Get()
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatalln(err)
	}

	moviesDoc := doc.Find(imdbTitleSelector)
	movies := make([]Movie, moviesDoc.Length())
	moviesDoc.Each(func(i int, s *goquery.Selection) {
		movies[i].Name = s.Text()
		movies[i].Link = s.AttrOr("href", "")
	})
	return movies
}

func (m *Movie) GetDetails() {
	req := imdbHttpClient()
	req.Url = fmt.Sprintf("%s%s", ImdbUrlBase, m.Link)
	response, err := req.Get()
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatalln(err)
	}

	creditSummaryLabelsDoc := doc.Find(creditSummaryLabelsSelector)
	creditSummaryLabelsDoc.Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Director:" || s.Text() == "Directors:" {
			creditSummaryValuesDoc := s.Parent().Find("a")
			m.Directors = make([]Artist, creditSummaryValuesDoc.Length())
			creditSummaryValuesDoc.Each(func(j int, s *goquery.Selection) {
				m.Directors[j].Name = s.Text()
				m.Directors[j].Link = s.AttrOr("href", "")
			})
		}
	})
}

func (m *Movie) Print() {
	fmt.Printf("%s (%s)\n", m.Name, m.Link)
	fmt.Println("\tDirector(s):")
	for _, director := range m.Directors {
		fmt.Printf("\t+\t%s (%s)\n", director.Name, director.Link)
	}
}

func main() {
	movies := findMovie(ExampleTitle)
	for _, movie := range movies {
		movie.GetDetails()
		movie.Print()
	}
}
