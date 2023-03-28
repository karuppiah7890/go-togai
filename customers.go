package togai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// TODO: Document all data types and fields.
// TODO: If one data type contains another
// data type, it should be linked so that
// the user can click it and see it, for
// example - Account field in Customer

// The primary use case of creating Customers and Accounts is to have the ability to track usage and revenue at the granular level.
//
// The Customers and Accounts are parent and child relationships. By default, Togai creates an Account for every Customer you create with the Customer ID and Customer Name same for the Account. You can also define the Account ID and Account Name by passing to the respective fields as given in the body parameters.
//
// You can use Customers and Accounts for Company and Users relationship, Customer and Multiple environments like staging, sandbox, and production, or any similar use case.
//
// Learn more from [Customer Management]
//
// [Customer Management]: https://docs.togai.com/docs/customers
type Customer struct {
	// Customer identifier
	Id string `json:"id"`
	// Name of the Customer
	Name string `json:"name"`
	// Primary email of the customer
	PrimaryEmail string `json:"primaryEmail"`
	// Billing address of the customer
	BillingAddress string `json:"billingAddress"`
	// Customer Settings
	Settings []Setting `json:"settings"`
	// Payload to create [Account]
	Account Account `json:"account"`
}

type Setting struct {
	Id        string          `json:"id"`
	Value     string          `json:"value"`
	Namespace Namespace       `json:"namespace"`
	Name      string          `json:"name"`
	DataType  SettingDataType `json:"dataType"`
}

type Namespace string

const (
	UserNamespace  Namespace = "USER"
	TogaiNamespace Namespace = "TOGAI"
)

type SettingDataType string

const (
	StringSettingDataType    SettingDataType = "STRING"
	NumericSettingDataType   SettingDataType = "NUMERIC"
	JsonSettingDataType      SettingDataType = "JSON"
	JsonLogicSettingDataType SettingDataType = "JSON_LOGIC"
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

type CreateCustomerOutput struct {
	// Customer identifier
	Id string `json:"id"`
	// Name of the Customer
	Name string `json:"name"`
	// Primary email of the customer
	PrimaryEmail string `json:"primaryEmail"`
	// Billing address of the customer
	BillingAddress string `json:"billingAddress"`
	// Customer Settings
	Settings []Setting `json:"settings"`
	// Payload to create [Account]
	Account CreateAccountOutput `json:"account"`
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

// TODO: Move URL and Token to togai client struct.
// Convert this function into a method on togai client struct and use url and token from there.
// Use http client from the togai client struct
func (c *TogaiClient) CreateCustomer(customer Customer) (*CreateCustomerOutput, error) {
	customersEndpoint, err := url.JoinPath(c.ApiBaseUrl, "customers")
	if err != nil {
		return nil, fmt.Errorf("error forming customers API endpoint: %v", err)
	}

	createCustomerJsonPayload, err := json.Marshal(customer)
	if err != nil {
		return nil, fmt.Errorf("error serializing customer object to JSON string: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, customersEndpoint, bytes.NewReader(createCustomerJsonPayload))
	if err != nil {
		return nil, fmt.Errorf("error occurred while forming request: %v", err)
	}

	req.Header.Add(ACCEPT_HTTP_HEADER, JSON_TYPE)
	req.Header.Add(CONTENT_TYPE_HTTP_HEADER, JSON_TYPE)
	req.Header.Add(AUTHORIZATION_HTTP_HEADER, fmt.Sprintf("Bearer %s", c.ApiToken))

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

	var createdCustomer CreateCustomerOutput

	err = json.Unmarshal(body, &createdCustomer)
	if err != nil {
		return nil, fmt.Errorf("error occurred while parsing json response body: %v\n\njson body: %v", err, string(body))
	}

	return &createdCustomer, nil
}
