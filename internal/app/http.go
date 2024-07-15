package app

import (
	adapterhttp "go-time-tracker/internal/adapter/http"
	"go-time-tracker/internal/config"
	"net/http"
	"net/url"
)

// newHttpClient create http client with base url from config
func (a *App) newHttpClient(cfg config.ThirdPartyService) (*adapterhttp.Client, error) {
	baseUrl, err := url.Parse(cfg.Url)
	if err != nil {
		return nil, err
	}
	client := adapterhttp.Client{
		Client:  &http.Client{},
		BaseURL: baseUrl,
	}
	return &client, nil
}
