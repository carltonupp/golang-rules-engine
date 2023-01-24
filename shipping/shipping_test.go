package shipping_test

import (
	"testing"
	"time"

	"github.com/carltonupp/golang-rules-engine/basket"
	"github.com/carltonupp/golang-rules-engine/customer"
	"github.com/carltonupp/golang-rules-engine/shipping"
)

func TestAddRule(t *testing.T) {
	calc := shipping.ShippingCalculator{}

	rule := shipping.NewRule("Overseas Standard Shipping", func(b basket.Basket) shipping.RuleRunResult {
		if b.Customer.IsOverseas {
			return shipping.RuleRunResult{
				Success: true,
				Value:   19.99,
			}
		}

		return shipping.RuleRunResult{
			Success: false,
			Value:   9.99,
		}
	})

	calc.AddRule(rule)

	if len(calc.Rules) != 1 {
		t.Errorf("shipping.AddRule failed. Expected: %d, Received: %d", 1, len(calc.Rules))
	}
}

func TestCalculateLowestShipping(t *testing.T) {

	cust := customer.Customer{
		Name:        "John Smith",
		DateOfBirth: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		IsOverseas:  true,
	}

	line := basket.NewLine("Spark Plugs", 19.99, 2)

	b := basket.New(cust)
	b.AddLine(line)
	calc := shipping.ShippingCalculator{}

	rule := shipping.NewRule("Overseas Standard Shipping", func(b basket.Basket) shipping.RuleRunResult {
		if b.Customer.IsOverseas {
			return shipping.RuleRunResult{
				Success: true,
				Value:   19.99,
			}
		}

		return shipping.RuleRunResult{
			Success: false,
			Value:   9.99,
		}
	})

	calc.AddRule(rule)
	result := calc.CalculateLowestShipping(b)

	if result != 19.99 {
		t.Errorf("shipping.CalculateLowestShipping failed. Expected: %f, Got: %f", 19.99, result)
	}
}
