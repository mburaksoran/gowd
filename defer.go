package main

import "fmt"

// defer
// kendisini çevreleyen fonksiyon dönene kadar fonksiyonun çalışmasını erteler.

var isConnected bool = false

func main() {
	fmt.Printf("Connection Open: %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connection Open: %v\n", isConnected)
}

func databaseProcessing() {
	connect()
	defer disconnect()
	defer disconnect2()
	// .....
	// .....
	// .....
	fmt.Println("Do something...")
	fmt.Println("Connection Open: end of scope")
}

func connect() {
	isConnected = true
	fmt.Println("Connected to database!")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected!")
}

func disconnect2() {
	isConnected = false
	fmt.Println("Disconnected! 2")
}
