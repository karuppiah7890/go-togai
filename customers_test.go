package togai_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/karuppiah7890/go-togai"
)

func TestCustomers(t *testing.T) {
	apiBaseUrl := os.Getenv("API_BASE_URL")
	apiToken := os.Getenv("API_TOKEN")
	c := togai.NewTogaiClient(apiBaseUrl, apiToken)

	t.Run("create customer", func(t *testing.T) {
		randomNumber := rand.Int()
		customer := togai.Customer{
			Id:           fmt.Sprintf("test-customer-%d", randomNumber),
			Name:         fmt.Sprintf("Test Customer %d", randomNumber),
			PrimaryEmail: "karuppiah+test-customer@togai.com",
			Settings: []togai.Setting{
				{
					Id:        fmt.Sprintf("dummy-customer-setting-%d", randomNumber),
					Name:      fmt.Sprintf("Dummy Customer Setting %d", randomNumber),
					Namespace: togai.UserNamespace,
					Value:     "10",
					DataType:  togai.NumericSettingDataType,
				},
			},
			BillingAddress: "Test Billing Address",
			Account: togai.Account{
				Id:              fmt.Sprintf("test-account-%d", randomNumber),
				Name:            fmt.Sprintf("Test Account %d", randomNumber),
				InvoiceCurrency: "USD",
				Aliases:         []string{fmt.Sprintf("test-account-%d", randomNumber)},
				Settings: []togai.Setting{
					{
						Id:        fmt.Sprintf("dummy-account-setting-%d", randomNumber),
						Name:      fmt.Sprintf("Dummy Account Setting %d", randomNumber),
						Namespace: togai.UserNamespace,
						Value:     fmt.Sprintf("something %d", randomNumber),
						DataType:  togai.StringSettingDataType,
					},
				},
			},
		}
		// TODO: Check if the output response fields and input request fields match, and that response also has
		// some expected values like status as ACTIVE for account and account alias
		_, err := c.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("expected no error but an error occurred: %v", err)
		}
	})
}
