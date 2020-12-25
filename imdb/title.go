package imdb

import (
	"log"
	"strings"

	"github.com/abhishekkr/imdb-cli/config"

	"github.com/PuerkitoBio/goquery"
)

type TitleCategory int

const (
	TcAll TitleCategory = iota
	TcMovie
	TcTV
	TcTVEpisode
	TcVideoGame
)

func (tc TitleCategory) getParam() string {
	return []string{
		"_", // any value for All Title works until conflicts with valid
		config.GetParamForMovieTitle,
		config.GetParamForTVTitle,
		config.GetParamForTVEpisodeTitle,
		config.GetParamForVideoGameTitle,
	}[tc]
}

func FindTitle(query string, category TitleCategory, exact bool) []Movie {
	req := imdbHttpClient()
	req.Url = config.FindUrlBase
	req.GetParams["s"] = "tt"
	req.GetParams["ttype"] = category.getParam()
	req.GetParams["ref_"] = "fn_" + req.GetParams["ttype"]
	req.GetParams["q"] = query
	if exact {
		req.GetParams["exact"] = "true"
	}
	response, err := req.Get()
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatalln(err)
	}

	moviesDoc := doc.Find(config.ImdbTitleSelector)
	movies := make([]Movie, moviesDoc.Length())
	moviesDoc.Each(func(i int, s *goquery.Selection) {
		movies[i].Name = s.Text()
		movies[i].Link = s.AttrOr("href", "")
	})
	return movies
}
