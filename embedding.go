package main

import "fmt"

func main() {
	// x := new(Account)
	// x.ID = 1
	// x.AvailableFunds()
	// x.ProcessPayment(20)

	// y := new(CreditAccount)
	// y.ID = 0

	// ------

	x := &CreditAccount{}
	x.AvailableFunds()
}

type Account struct {
	ID int
}

func (a *Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds")
	return 0
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment")
	return true
}

type CreditAccount struct {
	Account
}
