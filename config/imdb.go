package config

const (
	ImdbUrlBase = "https://www.imdb.com"
	FindUrlBase = "https://www.imdb.com/find"

	GetParamForMovieTitle     = "ft"
	GetParamForTVTitle        = "tv"
	GetParamForTVEpisodeTitle = "ep"
	GetParamForVideoGameTitle = "vg"

	ImdbTitleSelector           = ".findList .findResult .result_text a"
	CreditSummaryLabelsSelector = ".credit_summary_item h4"
)
