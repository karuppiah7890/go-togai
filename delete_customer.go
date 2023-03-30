package togai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Status struct {
	Success bool `json:"success"`
}

func (c *TogaiClient) DeleteCustomer(customerId string) error {
	deleteCustomerEndpoint := c.apiBaseUrl.JoinPath("customers", customerId)

	req, err := http.NewRequest(http.MethodDelete, deleteCustomerEndpoint.String(), nil)
	if err != nil {
		return fmt.Errorf("error occurred while forming request: %v", err)
	}

	req.Header.Add(ACCEPT_HTTP_HEADER, JSON_TYPE)
	req.Header.Add(AUTHORIZATION_HTTP_HEADER, fmt.Sprintf("Bearer %s", c.apiToken))

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error occurred while sending request: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return fmt.Errorf("expected 200 OK response but got: \nstatus code: %v, status: %v, body: %v", res.StatusCode, res.Status, string(body))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error occurred while reading response body: %v", err)
	}

	var deleteCustomerOutput Status

	err = json.Unmarshal(body, &deleteCustomerOutput)
	if err != nil {
		return fmt.Errorf("error occurred while parsing json response body: %v\n\njson body: %v", err, string(body))
	}

	if !deleteCustomerOutput.Success {
		return fmt.Errorf("deletion of the customer did not succeed")
	}

	return nil
}
