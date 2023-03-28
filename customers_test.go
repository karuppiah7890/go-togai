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
			Id:             fmt.Sprintf("test-customer-%d", randomNumber),
			Name:           fmt.Sprintf("Test Customer %d", randomNumber),
			PrimaryEmail:   "karuppiah+test-customer@togai.com",
			BillingAddress: "Test Billing Address",
			Account: togai.Account{
				Id:              fmt.Sprintf("test-account-%d", randomNumber),
				Name:            fmt.Sprintf("Test Account %d", randomNumber),
				InvoiceCurrency: "USD",
				Aliases:         []string{fmt.Sprintf("test-account-%d", randomNumber)},
			},
		}
		err := c.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("expected no error but an error occurred: %v", err)
		}
	})
}
