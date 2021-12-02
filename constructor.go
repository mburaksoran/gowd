package main

import "fmt"

func main() {
	x := new(User)
	x.FirstName = "C"
	x.LastName = "O"
	// result := GetFullName(x)
	// result := x.GetFullName()
	// fmt.Println(result)

	x.SetFirstName("CO")
	fmt.Println(x.FirstName)
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

// Function Version

// func GetFullName(user *User) string {
// return user.FirstName + " " + user.LastName
// }

// Method Version
// Method'lar bir nesneye aittir.

// func (u *User) GetFullName() string {
// return u.FirstName + " " + u.LastName
// }

// func (u User) GetFullName() string {
// return u.FirstName + " " + u.LastName
// }

// * kullanmadığımız versiyonda hafızadaki hangi referansın güncellendiğini bilemeyiz.
// En azından bizim nesne örneğimizin güncellenmeyeceği garantidir.
func (u *User) SetFirstName(firstname string) {
	u.FirstName = firstname
}
