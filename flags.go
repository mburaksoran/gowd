package main

import (
	"flag"
	"os"
	"strconv"
)

func main() {

	/* name := flag.String("name", "Gopher", "name of the gopher")
	age := flag.Int("age", 2, "age of the gopher")
	cute := flag.Bool("cute", false, "is the gopher cute?")
	flag.Parse()
	fmt.Printf("gopher stats\nName: %s\nAge: %d\nCute: %t\n", *name, *age, *cute)
	*/

	var (
		name string
		age  int
		cute bool
	)

	flag.StringVar(&name, "name", defaultName(), "name of the gopher")
	flag.IntVar(&age, "age", defaultAge(), "age of the gopher")
	flag.IntVar(&age, "age", defaultAge(), "age of the gopher")
	flag.BoolVar(&cute, "cute", defaultCuteness(), "is gopher cute?")
	flag.BoolVar(&cute, "cute", defaultCuteness(), "is gopher cute?")
	flag.Parse()
	/* fmt.Printf("Gopher stats\nName: %s\nAge: %d\nCute: %t\n", *name, *age, *cute) */
}

func defaultName() string {
	if os.Getenv("GOPHER_DEFAULT_NAME") != "" {
		return os.Getenv("GOPHER_DEFAULT_NAME")
	}
	return "Gopher"
}

func defaultAge() int {
	if os.Getenv("GOPHER_DEFAULT_AGE") != "" {
		age, err := strconv.Atoi(os.Getenv("GOPHER_DEFAULT_AGE"))
		if err == nil {
			return age

		}

	}
	return 7
}

func defaultCuteness() bool {
	if os.Getenv("GOPHER_DEFAULT_CUTENESS") != "" {
		cute, err := strconv.ParseBool(os.Getenv("GOPHER_DEFAULT_CUTENESS"))
		if err == nil {
			return cute
		}
	}
	return true
}

// go run flags.go -name=CO -age=33 -cute=false
