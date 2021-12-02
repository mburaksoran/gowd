package main

import (
	"fmt"
)

type level int

const (
	UNSPECIFIED level = iota // 0 ile başlar
	TRACE
	INFO
	WARNING
	ERROR
)

// köşeli parantez içerisine ne kadar veri alacağını belirtiyor. eğer boyut belli değilse ... ile ifade ediliyor
var levels = [...]string{
	"UNSPECIFIED",
	"TRACE",
	"INFO",
	"WARNING",
	"ERROR",
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
func main() {
	x := 44.9

	fmt.Println(int(x))

}
