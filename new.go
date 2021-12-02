package main

import "fmt"

func main() {
	x := new(User)
	x.FirstName = "C"
	x.LastName = "O"
	result := GetFullName(x)

	fmt.Println(result)
	fmt.Println(x.GetFullName())
}

type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Age       int
}

func NewUser() *User {
	return new(User)
}

// function version
func GetFullName(user *User) string {
	return user.FirstName + " " + user.LastName
}

// method version
func (r *User) GetFullName() string {
	return r.FirstName + " " + r.LastName
}
