package main

import "fmt"

func main() {

	// v1

	u1 := &User{
		ID:        5,
		FirstName: "Ali",
		LastName:  "Veli",
		UserName:  "aliveli",
		Age:       29,
		Pay: &Payment{
			Salary: 3.555,
			Bonus:  522,
		},
	}

	fmt.Println(u1.GetFullName())
	fmt.Println(u1.GetPayment())

}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
	Pay       *Payment
}

type Payment struct {
	ID     int
	Salary float64
	Bonus  float64
}

// Constructor

func NewUser() *User {
	u := new(User)
	u.Pay = NewPayment()
	return u
}

func NewPayment() *Payment {
	return new(Payment)
}

// Methods

func (u *User) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) SetUserName(uname string) {
	u.UserName = uname
}

func (u *User) GetPayment() float64 {
	return u.Pay.Salary + u.Pay.Bonus
}
