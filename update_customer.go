package togai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UpdateCustomerInput struct {
	Id string `json:"-"`
	// Name of the Customer
	Name *string `json:"name,omitempty"`
	// Primary email of the customer
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	// Billing address of the customer
	BillingAddress *string `json:"billingAddress,omitempty"`
}

// UpdateCustomer updates the given customer
func (c *TogaiClient) UpdateCustomer(customer UpdateCustomerInput) (*Customer, error) {
	updateCustomersEndpoint := c.apiBaseUrl.JoinPath("customers", customer.Id)

	updateCustomerJsonPayload, err := json.Marshal(customer)
	if err != nil {
		return nil, fmt.Errorf("error serializing customer object to JSON string: %v", err)
	}

	req, err := http.NewRequest(http.MethodPatch, updateCustomersEndpoint.String(), bytes.NewReader(updateCustomerJsonPayload))
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

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("expected 200 OK response but got: \nstatus code: %v, status: %v, body: %v", res.StatusCode, res.Status, string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error occurred while reading response body: %v", err)
	}

	var updatedCustomer Customer

	err = json.Unmarshal(body, &updatedCustomer)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing json response body: %v\n\njson body: %v", err, string(body))
	}

	return &updatedCustomer, nil
}
