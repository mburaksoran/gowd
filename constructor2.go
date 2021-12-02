package main

import "fmt"

func main() {
	x := new(User)
	x.FirstName = "Cihan"
	x.LastName = "Ozhan"

	// p := new(Payment)
	// p.Amount = 10.2
	// p.Currency = "USD"

	x.Pay.Amount = 10.3
	x.Pay.Currency = "USD"

	fmt.Println(x.Pay.Amount)
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       Payment
}

type Payment struct {
	ID       int
	Amount   float64
	Currency string
}

// Constructor

func NewUser() *User {
	return new(User)
}

// func NewPayment() *Payment {
// return new(Payment)
// }
