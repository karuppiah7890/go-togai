package togai

import "net/http"

type TogaiClient struct {
	ApiBaseUrl string
	ApiToken   string
	httpClient *http.Client
}

const (
	JSON_TYPE                 = "application/json"
	ACCEPT_HTTP_HEADER        = "accept"
	CONTENT_TYPE_HTTP_HEADER  = "content-type"
	AUTHORIZATION_HTTP_HEADER = "authorization"
)

func NewTogaiClient(ApiBaseUrl string, ApiToken string) *TogaiClient {
	return &TogaiClient{
		ApiBaseUrl: ApiBaseUrl,
		ApiToken:   ApiToken,
		httpClient: http.DefaultClient,
	}
}
