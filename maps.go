package main

import (
	"fmt"
	"sort"
)

// Map (KeyValuePair)

// map[int]string

func main() {

	states := make(map[string]string)
	fmt.Println("Boş: ", states)

	states["IST"] = "Istanbul"
	states["ANK"] = "Ankara"
	states["ANT"] = "Antalya"
	fmt.Println("Dolu: ", states)

	// antalya := states["ANT"]
	// fmt.Println("Seçili Şehir: ", antalya)

	delete(states, "ANK")
	// fmt.Println(states)

	states["ERZ"] = "Erzurum"
	// fmt.Println(states)

	// for k, v := range states {
	// fmt.Printf("%v: %v\n", k, v)
	// }

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	languages := map[string]string{
		"TR": "Turkish",
		"EN": "English",
		"ES": "Spanish",
	}

	filter := "TR"
	if lang, ok := languages[filter]; ok {
		fmt.Println(lang)
	} else {
		fmt.Println("Key doesn't exist.")
	}

	removeIt(languages, "ES")
	for _, v := range languages {
		fmt.Printf(v + "\n")
	}

	sort.Strings(keys)
}

func removeIt(langs map[string]string, key string) {
	delete(langs, key)
}
