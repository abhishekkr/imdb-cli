package config

const (
	ImdbUrlBase          = "https://www.imdb.com"
	FindUrlBase          = "https://www.imdb.com/find"
	FullCreditsUrlSuffix = "/fullcredits"

	GetParamForMovieTitle     = "ft"
	GetParamForTVTitle        = "tv"
	GetParamForTVEpisodeTitle = "ep"
	GetParamForVideoGameTitle = "vg"

	ImdbItemSelector  = ".ipc-metadata-list-summary-item"
	ImdbTitleSelector = ".ipc-title__text"
	ImdbLinkSelector  = ".ipc-title-link-wrapper"

	CreditSectionSelector = ".ipc-page-section"
	CreditLabelSelector   = ".ipc-title__text"
	CreditValuesSelector  = ".name-credits--title-text-big"
)

var (
	FullCreditsUrlParams = map[string]string{"ref_": "tt_cl_sm#cast"}
)
