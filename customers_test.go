package togai_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/karuppiah7890/go-togai"
)

// TODO: Delete test data that we create after the tests are done.
// Deleting test data would use the APIs we test as part of the
// library tests, so, let's hope that the library works and hence
// the cleanup done with the library also works when doing the testing

func TestCustomers(t *testing.T) {
	apiBaseUrl := os.Getenv("API_BASE_URL")
	apiToken := os.Getenv("API_TOKEN")
	c, err := togai.NewTogaiClient(apiBaseUrl, apiToken)
	if err != nil {
		t.Fatalf("expected no error while creating togai client but an error occurred: %v", err)
	}

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
			t.Fatalf("expected no error while creating customer but an error occurred: %v", err)
		}
	})

	t.Run("list customers", func(t *testing.T) {
		// TODO: Check if the output response fields and input request fields match, and that response also has
		// some expected values like status as ACTIVE for account and account alias
		numberOfCustomers := 2

		customers, err := c.ListCustomers("", numberOfCustomers)
		if err != nil {
			t.Fatalf("expected no error while listing customers but an error occurred: %v", err)
		}
		if len(customers.Data) != numberOfCustomers {
			t.Fatalf("expected %d customers while listing customers but there's %d customers", numberOfCustomers, len(customers.Data))
		}
	})

}
