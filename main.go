package main

import (
	"flag"
	"log"

	imdb "github.com/abhishekkr/imdb-cli/imdb"
)

var (
	Title      = flag.String("title", "", "title to search for")
	ExactMatch = flag.Bool("exact", false, "title to search for")
)

func main() {
	flag.Parse()
	if *Title != "" {
		log.Printf("looking for: %v", *Title)
		movies := imdb.FindMovie(*Title, *ExactMatch)
		for _, movie := range movies {
			movie.GetDetails()
			movie.Print()
		}
	} else {
		log.Println("No param passed for any task, try '-help' to see capabilities.")
	}
}
