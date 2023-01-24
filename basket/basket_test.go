package basket_test

import (
	"testing"
	"time"

	"github.com/carltonupp/golang-rules-engine/basket"
	"github.com/carltonupp/golang-rules-engine/customer"
)

func TestNewBasket(t *testing.T) {
	expected := basket.Basket{
		Customer: customer.Customer{
			Name:        "John Smith",
			DateOfBirth: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			IsOverseas:  true,
		},
	}

	cust := customer.Customer{
		Name:        "John Smith",
		DateOfBirth: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsOverseas:  true,
	}

	result := basket.New(cust)

	if expected.Customer != result.Customer {
		t.Errorf("basket.New failed. Expected: %+v\n, Got: %+v\n", expected, result)
	}
}

func TestNewLine(t *testing.T) {
	expected := basket.Line{
		Description: "Spark Plugs",
		UnitPrice:   19.99,
		Quantity:    2,
	}

	result := basket.NewLine("Spark Plugs", 19.99, 2)

	if expected != result {
		t.Errorf("basket.NewLine failed. Expected: %+v\n, Got: %+v\n", expected, result)
	}
}

func TestAddLine(t *testing.T) {

	cust := customer.Customer{
		Name:        "John Smith",
		DateOfBirth: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsOverseas:  true,
	}

	line := basket.NewLine("Spark Plugs", 19.99, 2)

	expected := basket.Basket{
		Customer: cust,
		Lines: []basket.Line{
			line,
		},
	}

	basket := basket.New(cust)
	basket.AddLine(line)

	if len(expected.Lines) != len(basket.Lines) {
		t.Errorf("basket.AddLine failed. Expected: %+v\n, Got: %+v\n", expected, basket)
	}
}

func TestGetBasketTotal(t *testing.T) {
	input := basket.Basket{
		Customer: customer.Customer{
			Name:        "John Smith",
			DateOfBirth: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
			IsOverseas:  true,
		},
		Lines: []basket.Line{
			{
				Description: "Spark Plugs",
				UnitPrice:   19.99,
				Quantity:    2,
			},
		},
	}

	result := input.GetBasketTotal()
	expected := 39.98

	if expected != result {
		t.Errorf("basket.GetBasketTotal failed. Expected: %f, Got: %f", expected, result)
	}
}
