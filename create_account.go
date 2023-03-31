package togai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Account struct {
	// Identifier of the account
	Id string `json:"id"`
	// Name of the Account
	Name string `json:"name"`
	// Use [ISO 4217] currency code in which the account must be invoiced.
	// For example: AED is the currency code for United Arab Emirates dirham.
	//
	// [ISO 4217]: https://en.wikipedia.org/wiki/ISO_4217
	InvoiceCurrency string `json:"invoiceCurrency"`
	// Aliases are tags that are associated with an account. Multiple aliases are allowed for a single account.
	Aliases []string `json:"aliases"`
	// Account Settings
	Settings []Setting `json:"settings"`
}

type CreateAccountOutput struct {
	// Identifier of the account
	Id string `json:"id"`
	// Name of the Account
	Name string `json:"name"`
	// Status of the account
	Status string `json:"status"`
	// Use [ISO 4217] currency code in which the account must be invoiced.
	// For example: AED is the currency code for United Arab Emirates dirham.
	//
	// [ISO 4217]: https://en.wikipedia.org/wiki/ISO_4217
	InvoiceCurrency string `json:"invoiceCurrency"`
	// Aliases are tags that are associated with an account. Multiple aliases are allowed for a single account.
	Aliases []AccountAlias `json:"aliases"`
	// Account Settings
	Settings []Setting `json:"settings"`
}

type AccountAlias struct {
	Alias  string `json:"alias"`
	Status string `json:"status"`
}

// CreateAccount creates the given account
func (c *TogaiClient) CreateAccount(customerId string, account Account) (*CreateAccountOutput, error) {
	createAccountEndpoint := c.apiBaseUrl.JoinPath("customers", customerId, "accounts")

	createAccountJsonPayload, err := json.Marshal(account)
	if err != nil {
		return nil, fmt.Errorf("error serializing customer object to JSON string: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, createAccountEndpoint.String(), bytes.NewReader(createAccountJsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error occurred while forming request: %v", err)
	}

	req.Header.Add(ACCEPT_HTTP_HEADER, JSON_TYPE)
	req.Header.Add(CONTENT_TYPE_HTTP_HEADER, JSON_TYPE)
	req.Header.Add(AUTHORIZATION_HTTP_HEADER, fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error occurred while sending request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 201 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("expected 201 created response but got: \nstatus code: %v, status: %v, body: %v", res.StatusCode, res.Status, string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred while reading response body: %v", err)
	}

	var createdAccount CreateAccountOutput

	err = json.Unmarshal(body, &createdAccount)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing json response body: %v\n\njson body: %v", err, string(body))
	}

	return &createdAccount, nil
}
