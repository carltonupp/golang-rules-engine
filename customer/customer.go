package customer

import "time"

type Customer struct {
	Name        string
	DateOfBirth time.Time
	IsOverseas  bool
}

func New(name string, dob time.Time, isOverseas bool) Customer {
	return Customer{
		Name:        name,
		DateOfBirth: dob,
		IsOverseas:  isOverseas,
	}
}
