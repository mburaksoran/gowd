package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2021, time.November, 10, 20, 0, 0, 0, time.UTC)

	fmt.Printf("Go Launch at %s\n", t)

	now := time.Now()

	fmt.Printf("The time noww is %s\n", now)

	fmt.Println("The month is ", t.Month())

	fmt.Println("the day is", t.Day())

	fmt.Println("the weekday is", t.Weekday())

	fmt.Println("----------")

	// tarihe 1 gün ekleme

	tomorrow := t.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	longFormat := "Monday, January 2, 2006"

	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

}
