package imdb

import (
	"github.com/abhishekkr/gol/golhttpclient"
	"github.com/abhishekkr/imdb-cli/config"
)

func imdbHttpClient() *golhttpclient.HTTPRequest {
	req := golhttpclient.HTTPRequest{
		GetParams: map[string]string{},
		HTTPHeaders: map[string]string{
			"user-agent": config.HTTPUserAgent,
		},
	}
	golhttpclient.SkipSSLVerify = true
	return &req
}
