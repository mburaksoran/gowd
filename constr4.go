package main

import "fmt"

func main() {

	// v2

	// User Creation
	user := NewUser()
	user.FirstName = "Cihan"
	user.LastName = "Özhan"
	user.Age = 33
	user.UserName = "cihanozhan"

	// Payment Creation v1

	// user.Pay.Salary = 50
	// user.Pay.Bonus = 5
	// // Printing
	// fmt.Println(user.Pay)
	// fmt.Println(user.GetPayment())

	// Payment Creation v2

	user.Pay = &Payment{Salary: 45, Bonus: 3}
	fmt.Println(user.GetFullName())
	// fmt.Println(user.Pay)
	fmt.Println(user.GetPayment())

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
