package main

import "fmt"

//map[int]string

func main() {
	/*
		scores := [3]int{}
		scores[0] = 32
		scores[1] = 33
		scores[2] = 44
		fmt.Println(scores)

		var scores2 [3]int
		scores2[0] = 32
		scores2[1] = 33
		scores2[2] = 44
		fmt.Println(scores2)

		scores3 := []int{22, 33, 44, 55}
		scores3[0] = 32
		scores3[1] = 33
		scores3[2] = 44
		fmt.Println(scores3)
		fmt.Println(len(scores3))
		scores3 = append(scores3, 24)
		fmt.Println(len(scores3))
		// variadic function */

	/* 	var cars3 [3]string
	   	cars3[0] = "mercedes"
	   	cars3[0] = "Tesla"
	   	cars3[0] = "volvo"

	   	i := 0
	   	for i <= len(cars3)-1 {
	   		fmt.Println(cars3[i])
	   	}

	   	for j := 0; j < len(cars3); j++ {

	   		fmt.Println(cars3[j])
	   	}

	   	for k, v := range cars3 {

	   		fmt.Println("K=", k, "Value=", v)

	   	} */

	myArrays1 := [...]int{42, 27, 99}

	mySlice1 := myArrays1[:]

	fmt.Println(mySlice1)

	categories := make(map[int]string)
	categories[1] = "Java"
	categories[2] = "python"
	categories[3] = "R"
	categories[4] = "golang"

	fmt.Println(categories[0])

	states := make(map[string]string)
	fmt.Println(states)
	states["IST"] = "istanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println(states["ANK"])

	delete(states, "ANK")

	//Ekleme
	states["ERZ"] = "Erzurum"

}
