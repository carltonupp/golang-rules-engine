package customer_test

import (
	"testing"
	"time"

	"github.com/carltonupp/golang-rules-engine/customer"
)

func TestNewCustomer(t *testing.T) {
	expected := customer.Customer{
		Name:        "John Smith",
		DateOfBirth: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsOverseas:  true,
	}

	dob := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	result := customer.New("John Smith", dob, true)

	if result != expected {
		t.Errorf("Test failed. Expected: %+v\n, Received: %+v\n", expected, result)
	}
}
