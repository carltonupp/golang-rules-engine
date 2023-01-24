package shipping

import "github.com/carltonupp/golang-rules-engine/basket"

type RuleRunResult struct {
	Success bool
	Value   float64
}

type Rule struct {
	Name      string
	Condition func(basket.Basket) RuleRunResult
	Output    float64
}

func NewRule(name string, condition func(basket.Basket) RuleRunResult) Rule {
	return Rule{
		Name:      name,
		Condition: condition,
	}
}
