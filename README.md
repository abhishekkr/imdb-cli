
## imdb-cli

> a cli utility trying to provide few IMDB.com info on console

* Build `go build -o out/imdb main.go` or `./build.sh`


* Search for all Titles matching a term

> * for general search of title `./out/imdb -title "Fight Club"`
>
> * to search for exact match `./out/imdb -title "Fight Club" -exact`; using the switch `-exact`


* To restrict query for just movie use `-movie`, `-tv` for TV, `-episode` for TV Episode and `-game` for Video Game. By default it will search for **all** titles.

> * for general search of movie `./out/imdb -title "Fight Club" -movie`

---
