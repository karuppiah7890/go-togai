package togai_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/karuppiah7890/go-togai"
)

func TestAccounts(t *testing.T) {
	apiBaseUrl := os.Getenv("API_BASE_URL")
	apiToken := os.Getenv("API_TOKEN")
	c, err := togai.NewTogaiClient(apiBaseUrl, apiToken)
	if err != nil {
		t.Fatalf("expected no error while creating togai client but an error occurred: %v", err)
	}
	t.Run("create account", func(t *testing.T) {
		randomNumber := rand.Int()
		customer := dummyCustomerWithAccount(randomNumber)
		// TODO: Check if the output response fields and input request fields match, and that response also has
		// some expected values like status as ACTIVE for account and account alias
		_, err := c.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("expected no error while creating customer but an error occurred: %v", err)
		}

		anotherRandomNumber := rand.Int()
		account := dummyAccount(anotherRandomNumber)
		// TODO: Check if the output response fields and input request fields match, and that response also has
		// some expected values like status as ACTIVE for account and account alias
		_, err = c.CreateAccount(customer.Id, account)
		if err != nil {
			t.Fatalf("expected no error while creating account but an error occurred: %v", err)
		}
	})

	t.Run("delete account", func(t *testing.T) {
		randomNumber := rand.Int()
		customer := dummyCustomerWithAccount(randomNumber)
		// TODO: Check if the output response fields and input request fields match, and that response also has
		// some expected values like status as ACTIVE for account and account alias
		_, err := c.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("expected no error while creating customer but an error occurred: %v", err)
		}

		anotherRandomNumber := rand.Int()
		account := dummyAccount(anotherRandomNumber)
		// TODO: Check if the output response fields and input request fields match, and that response also has
		// some expected values like status as ACTIVE for account and account alias
		_, err = c.CreateAccount(customer.Id, account)
		if err != nil {
			t.Fatalf("expected no error while creating account but an error occurred: %v", err)
		}
	})

}

func dummyAccount(randomNumber int) togai.Account {
	return togai.Account{
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
	}
}
