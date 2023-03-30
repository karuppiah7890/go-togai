package togai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ListCustomersOutput struct {
	Data      []Customer `json:"data"`
	NextToken string     `json:"nextToken"`
}

type Customer struct {
	// Customer identifier
	Id string `json:"id"`
	// Name of the Customer
	Name string `json:"name"`
	// Primary email of the customer
	PrimaryEmail string `json:"primaryEmail"`
	// Billing address of the customer
	BillingAddress string `json:"billingAddress"`
	// Status of the customer
	Status string `json:"status"`
	// Time at which the customer was created
	CreatedAt time.Time `json:"createdAt"`
	// Time at which the customer was last updated
	UpdatedAt time.Time `json:"updatedAt"`
}

const (
	NEXT_TOKEN_QUERY_PARAM = "nextToken"
	PAGE_SIZE_QUERY_PARAM  = "pageSize"
)

func (c *TogaiClient) ListCustomers(nextToken string, pageSize int) (*ListCustomersOutput, error) {
	listCustomersEndpoint := c.apiBaseUrl.JoinPath("customers")

	queryValues := listCustomersEndpoint.Query()

	if strings.TrimSpace(nextToken) != "" {
		queryValues.Set(NEXT_TOKEN_QUERY_PARAM, nextToken)
	}

	if pageSize != 0 {
		queryValues.Set(PAGE_SIZE_QUERY_PARAM, strconv.Itoa(pageSize))
	}

	listCustomersEndpoint.RawQuery = queryValues.Encode()

	req, err := http.NewRequest(http.MethodGet, listCustomersEndpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error occurred while forming request: %v", err)
	}

	req.Header.Add(ACCEPT_HTTP_HEADER, JSON_TYPE)
	req.Header.Add(AUTHORIZATION_HTTP_HEADER, fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occurred while sending request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("expected 200 OK response but got: \nstatus code: %v, status: %v, body: %v", res.StatusCode, res.Status, string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred while reading response body: %v", err)
	}

	var listCustomersOutput ListCustomersOutput

	err = json.Unmarshal(body, &listCustomersOutput)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing json response body: %v\n\njson body: %v", err, string(body))
	}

	return &listCustomersOutput, nil
}
