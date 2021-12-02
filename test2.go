package main

import (
	"fmt"
)

func main() {
	var a, p int
	a, p = rectangle(20, 30)
	fmt.Println("Area:", a)
	fmt.Println("Parameter:", p)
	m, n := add(1, 2, 3, 454, 5)
	println(m, n)
}

func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func add(terms ...int) (int, int) {

	result := 0

	for _, term := range terms {
		result += term
	}
	return result, len(terms)
}

/* func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y

} */

//named result
func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y

}
