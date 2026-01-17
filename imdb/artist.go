package imdb

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Artist struct {
	Name string
	Link string
}

func queryArtists(s *goquery.Selection) []Artist {
	creditSummaryValuesDoc := s.Find(".name a")
	artists := make([]Artist, creditSummaryValuesDoc.Length())
	creditSummaryValuesDoc.Each(func(idx int, s *goquery.Selection) {
		artists[idx].Name = strings.TrimSpace(s.Text())
		artists[idx].Link = s.AttrOr("href", "")
	})
	return artists
}
