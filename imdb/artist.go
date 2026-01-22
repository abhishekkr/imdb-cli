package imdb

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/abhishekkr/imdb-cli/config"
)

type Artist struct {
	Name string
	Link string
}

func queryArtists(s *goquery.Selection) []Artist {
	creditValues := s.Find(config.CreditValuesSelector)
	artists := make([]Artist, creditValues.Length())
	creditValues.Each(func(idx int, s *goquery.Selection) {
		artists[idx].Name = strings.TrimSpace(s.Text())
		artists[idx].Link = s.AttrOr("href", "")
	})
	return artists
}
