package basket

import "github.com/carltonupp/golang-rules-engine/customer"

type Line struct {
	Description string
	UnitPrice   float64
	Quantity    int
}

type Basket struct {
	Customer customer.Customer
	Lines    []Line
}

func New(c customer.Customer) Basket {
	return Basket{
		Customer: c,
	}
}

func NewLine(desc string, unitPrice float64, qty int) Line {
	return Line{
		Description: desc,
		UnitPrice:   unitPrice,
		Quantity:    qty,
	}
}

func (b *Basket) AddLine(l Line) {
	b.Lines = append(b.Lines, l)
}

func (bd Basket) GetBasketTotal() float64 {
	var total float64
	for _, l := range bd.Lines {
		lt := l.UnitPrice * float64(l.Quantity)
		total += lt
	}
	return total
}
