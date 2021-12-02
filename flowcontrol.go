package main

import (
	"fmt"
	"strconv"
)

func main() {
	isSunny := true
	if isSunny {
		fmt.Println("Enjoy your day!")
	} else {
		fmt.Println("Tomorrow is another day!")
	}

	isSaturday := true

	if isSunny && isSaturday {

		fmt.Println("Let's swim outdoor.")

	}

	age := 15 // < >= <=

	if age > 13 {
		fmt.Println("yaşınız büyük olduğu için")
	} else {
		fmt.Println("yaşınız küçük")
	}

	// Nested if döngüleri if ile başlayıp else ile bitmeli.

	foo := 5

	if foo := 2; foo == 1 {
		fmt.Println("bar")
	} else {
		println("İç foo: " + "buz")
		println("iç foo: " + strconv.Itoa(foo))
	}
	println("Dış foo: " + strconv.Itoa(foo))

	colors := []string{"red", "green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
}
