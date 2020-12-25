package imdb

import (
	"fmt"
	"log"
	"strings"

	"github.com/abhishekkr/imdb-cli/config"

	"github.com/PuerkitoBio/goquery"
)

type Movie struct {
	Name      string
	Link      string
	Directors []Artist
}

func (m *Movie) GetDetails() {
	req := imdbHttpClient()
	req.Url = fmt.Sprintf("%s%s", config.ImdbUrlBase, m.Link)
	response, err := req.Get()
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatalln(err)
	}

	creditSummaryLabelsDoc := doc.Find(config.CreditSummaryLabelsSelector)
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
