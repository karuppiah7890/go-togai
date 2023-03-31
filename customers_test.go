package togai_test

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/karuppiah7890/go-togai"
	"github.com/stretchr/testify/assert"
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
		customer := dummyCustomerWithAccount(randomNumber)
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

	t.Run("get a customer", func(t *testing.T) {
		randomNumber := rand.Int()
		customer := dummyCustomerWithAccount(randomNumber)
		_, err := c.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("expected no error while creating customer but an error occurred: %v", err)
		}

		expectedCustomer := dummyCustomer(randomNumber)

		actualCustomer, err := c.GetCustomer(customer.Id)
		if err != nil {
			t.Fatalf("expected no error while getting a customer but an error occurred: %v", err)
		}
		assertCustomer(t, expectedCustomer, *actualCustomer)
	})

	t.Run("update a customer", func(t *testing.T) {
		randomNumber := rand.Int()
		customer := dummyCustomerWithAccount(randomNumber)
		_, err := c.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("expected no error while creating customer but an error occurred: %v", err)
		}

		expectedUpdatedCustomer := dummyCustomer(randomNumber)

		updatedName := customer.Name + " Updated"
		updatedPrimaryEmail := "updated-" + customer.PrimaryEmail
		updatedBillingAddress := "Updated " + customer.BillingAddress

		expectedUpdatedCustomer.Name = updatedName
		expectedUpdatedCustomer.PrimaryEmail = updatedPrimaryEmail
		expectedUpdatedCustomer.BillingAddress = updatedBillingAddress

		updateCustomerInput := togai.UpdateCustomerInput{
			Id:             customer.Id,
			Name:           &updatedName,
			PrimaryEmail:   &updatedPrimaryEmail,
			BillingAddress: &updatedBillingAddress,
		}

		updatedCustomer, err := c.UpdateCustomer(updateCustomerInput)
		if err != nil {
			t.Fatalf("expected no error while updating customer but an error occurred: %v", err)
		}
		assert.Equal(t, expectedUpdatedCustomer.Id, updatedCustomer.Id, "customer IDs should be equal")
		assert.Equal(t, expectedUpdatedCustomer.Name, updatedCustomer.Name, "customer names should be equal")
		assert.Equal(t, expectedUpdatedCustomer.PrimaryEmail, updatedCustomer.PrimaryEmail, "customer primary emails should be equal")
		assert.Equal(t, expectedUpdatedCustomer.BillingAddress, updatedCustomer.BillingAddress, "customer billing addresses should be equal")
	})

	t.Run("delete a customer", func(t *testing.T) {
		randomNumber := rand.Int()
		customer := dummyCustomerWithAccount(randomNumber)
		_, err := c.CreateCustomer(customer)
		if err != nil {
			t.Fatalf("expected no error while creating customer but an error occurred: %v", err)
		}

		err = c.DeleteAccount(customer.Id, customer.Account.Id)
		if err != nil {
			t.Fatalf("expected no error while deleting account but an error occurred: %v", err)
		}

		err = c.DeleteCustomer(customer.Id)
		if err != nil {
			t.Fatalf("expected no error while deleting customer but an error occurred: %v", err)
		}
	})
}

func dummyCustomerWithAccount(randomNumber int) togai.CustomerWithAccount {
	customer := dummyCustomer(randomNumber)
	return togai.CustomerWithAccount{
		Id:           customer.Id,
		Name:         customer.Name,
		PrimaryEmail: customer.PrimaryEmail,
		Settings: []togai.Setting{
			{
				Id:        fmt.Sprintf("dummy-customer-setting-%d", randomNumber),
				Name:      fmt.Sprintf("Dummy Customer Setting %d", randomNumber),
				Namespace: togai.UserNamespace,
				Value:     "10",
				DataType:  togai.NumericSettingDataType,
			},
		},
		BillingAddress: customer.BillingAddress,
		Account:        dummyAccount(randomNumber),
	}
}

func dummyCustomer(randomNumber int) togai.Customer {
	return togai.Customer{
		Id:             fmt.Sprintf("test-customer-%d", randomNumber),
		Name:           fmt.Sprintf("Test Customer %d", randomNumber),
		PrimaryEmail:   fmt.Sprintf("test-customer-%d@togai.com", randomNumber),
		BillingAddress: "Test Billing Address",
	}
}

func assertCustomer(t *testing.T, expectedCustomer togai.Customer, actualCustomer togai.Customer) {
	assert.Equal(t, expectedCustomer.Id, actualCustomer.Id, "customer IDs should be equal")
	assert.Equal(t, expectedCustomer.Name, actualCustomer.Name, "customer names should be equal")
	assert.Equal(t, expectedCustomer.PrimaryEmail, actualCustomer.PrimaryEmail, "customer primary emails should be equal")
	assert.Equal(t, expectedCustomer.BillingAddress, actualCustomer.BillingAddress, "customer billing addresses should be equal")
}
