package imdb

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/abhishekkr/imdb-cli/config"

	"github.com/PuerkitoBio/goquery"
)

type Title struct {
	Name       string
	Link       string
	Metadata   string
	Metascore  int
	IMDBRating string
	HasCredits bool
	Credits    Credits
}

func (tytl *Title) GetDetails() {
	req := imdbHttpClient()
	req.Url = fmt.Sprintf("%s%s%s", config.ImdbUrlBase, tytl.Link, config.FullCreditsUrlSuffix)
	req.GetParams = config.FullCreditsUrlParams
	response, err := req.Get()
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatalln(err)
	}
	tytl.Credits.get(doc)
}

func (tytl *Title) Print() {
	fmt.Printf("%s (%s)\n", tytl.Name, tytl.Link)
	if tytl.Metadata == "" {
		fmt.Println("\tProbably Awaiting Release")
	} else {
		fmt.Printf("\t%s\n", tytl.Metadata)
	}
	if tytl.Metascore >= 0 {
		fmt.Printf("\tMetascore: %d\n", tytl.Metascore)
	}
	if tytl.IMDBRating != "" {
		fmt.Printf("\tRating: %s\n", tytl.IMDBRating)
	}
	if !tytl.HasCredits {
		return
	}

	printArtists := func(title string, artists []Artist) {
		fmt.Printf("\t%s:\n", title)
		for _, artist := range artists {
			fmt.Printf("\t* %s (%s)\n", artist.Name, artist.Link)
		}
	}
	printArtists("Director(s)", tytl.Credits.Directors)
	printArtists("Writer(s)", tytl.Credits.Writers)
	printArtists("Cast", tytl.Credits.Cast)
	printArtists("Producer(s)", tytl.Credits.Producers)
	printArtists("Composer(s)", tytl.Credits.Composers)
	printArtists("Cinematographer(s)", tytl.Credits.Cinematographers)
}

func titleGetParams(query string, category TitleCategory, exact bool) map[string]string {
	ttype := category.httpGetParam()
	params := map[string]string{
		"s":     "tt",
		"ttype": ttype,
		"ref_":  "fn_" + ttype,
		"q":     query,
	}
	if exact {
		params["exact"] = "true"
	}
	return params
}

func FindTitle(query string, category TitleCategory, exact bool) []Title {
	req := imdbHttpClient()
	req.Url = config.FindUrlBase
	req.GetParams = titleGetParams(query, category, exact)
	response, err := req.Get()
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		log.Fatalln(err)
	}

	tytlLinkRe := regexp.MustCompile("/title/\\w+")

	tytlsDoc := doc.Find(config.ImdbItemSelector)
	tytls := make([]Title, tytlsDoc.Length())
	tytlsDoc.Each(func(i int, s *goquery.Selection) {
		tytls[i].Name = strings.TrimSpace(s.Find(config.ImdbTitleSelector).Text())
		tytlLink := s.Find(config.ImdbLinkSelector).AttrOr("href", "")
		tytls[i].Link = tytlLinkRe.FindStringSubmatch(tytlLink)[0]
		var metadata []string
		s.Find(".cli-title-metadata-item").Each(func(_ int, smeta *goquery.Selection) {
			metadata = append(metadata, smeta.Text())
		})

		tytls[i].Metadata = strings.TrimSpace(strings.Join(metadata, ", "))
		tytls[i].Metascore, err = strconv.Atoi(s.Find(".metacritic-score-box").Text())
		if err != nil {
			tytls[i].Metascore = -1
		}
		tytls[i].IMDBRating = strings.TrimSpace(s.Find(".ipc-rating-star--rating").Text() + " " + s.Find(".ipc-rating-star--voteCount").Text())
	})

	return tytls
}
