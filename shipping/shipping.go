package shipping

import (
	"github.com/carltonupp/golang-rules-engine/basket"
)

type ShippingCalculator struct {
	Rules []Rule
}

func (sc *ShippingCalculator) AddRule(r Rule) {
	sc.Rules = append(sc.Rules, r)
}

func (sc *ShippingCalculator) CalculateLowestShipping(b basket.Basket) float64 {
	var result float64
	for i, rule := range sc.Rules {
		val := rule.Condition(b)
		if i == 0 || val.Value < result {
			result = val.Value
		}
	}
	return result
}
