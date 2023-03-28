package togai

import (
	"fmt"
	"net/http"
	"net/url"
)

type TogaiClient struct {
	// TODO: Should these fields be exported? Maybe not, all fields can be internal
	apiBaseUrl *url.URL
	apiToken   string
	httpClient *http.Client
}

const (
	JSON_TYPE                 = "application/json"
	ACCEPT_HTTP_HEADER        = "accept"
	CONTENT_TYPE_HTTP_HEADER  = "content-type"
	AUTHORIZATION_HTTP_HEADER = "authorization"
)

func NewTogaiClient(apiBaseUrl string, apiToken string) (*TogaiClient, error) {
	parsedUrl, err := url.Parse(apiBaseUrl)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing api base url %s: %v", apiBaseUrl, err)
	}

	return &TogaiClient{
		apiBaseUrl: parsedUrl,
		apiToken:   apiToken,
		httpClient: http.DefaultClient,
	}, nil
}
