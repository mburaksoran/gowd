package main

// Struct

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	// v1

	// humX := Human{FirstName: "Murtaza", LastName: "Soylu"}
	// fmt.Println("Data : " + humX.FirstName)
	// fmt.Println("Data : " + humX.LastName)

	// humX := Human{FirstName: "Murtaza", LastName: "Soylu"}
	// humX := Human{"Murtaza", "Soylu", 54}

	// humX := Human{}
	// humX.FirstName = "Veli"
	// fmt.Println("Data : " + humX.FirstName)
	// fmt.Println("Data : " + humX.LastName)

	// v2

	// humY := new(Human)
	// humY.FirstName = "Ayşecan"
	// humY.LastName = "X"
	// fmt.Println("Data " + humY.FirstName)

	// v3

	// xx := NewHuman("CO", "Ozhan")
	// fmt.Println(xx.FirstName + " " + xx.LastName)

}

func NewHuman() *Human {
	return new(Human)
}

// func NewHuman(firstName, lastName string) *Human {
// x := new(Human)
// x.FirstName = firstName
// x.LastName = lastName
// return x
// }
