package main

import (
	"flag"
	"log"

	imdb "github.com/abhishekkr/imdb-cli/imdb"
)

var (
	flgTitle      = flag.String("title", "", "Title to search for")
	flgMovie      = flag.Bool("movie", false, "Movie to search for")
	flgTV         = flag.Bool("tv", false, "TV to search for")
	flgTVEpisode  = flag.Bool("episode", false, "TV Episode to search for")
	flgVideoGame  = flag.Bool("game", false, "Video Game to search for")
	flgExactMatch = flag.Bool("exact", false, "Exact Match to search for")
)

func titleCategory() imdb.TitleCategory {
	if *flgMovie {
		return imdb.TcMovie
	} else if *flgTV {
		return imdb.TcTV
	} else if *flgTVEpisode {
		return imdb.TcTVEpisode
	} else if *flgVideoGame {
		return imdb.TcVideoGame
	}
	return imdb.TcAll
}

func main() {
	flag.Parse()
	if *flgTitle != "" {
		log.Printf("looking for: %v", *flgTitle)
		movies := imdb.FindTitle(*flgTitle, titleCategory(), *flgExactMatch)
		for _, movie := range movies {
			movie.GetDetails()
			movie.Print()
		}
	} else {
		log.Println("No param passed for any task, try '-help' to see capabilities.")
	}
}
