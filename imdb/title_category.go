package imdb

import "github.com/abhishekkr/imdb-cli/config"

type TitleCategory int

const (
	TcAll TitleCategory = iota
	TcMovie
	TcTV
	TcTVEpisode
	TcVideoGame
)

func (tc TitleCategory) httpGetParam() string {
	return []string{
		"_", // any value for All Title works until conflicts with valid
		config.GetParamForMovieTitle,
		config.GetParamForTVTitle,
		config.GetParamForTVEpisodeTitle,
		config.GetParamForVideoGameTitle,
	}[tc]
}
