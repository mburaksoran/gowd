package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var isExist bool = false
var cus Customer
var customers []Customer

func main() {
	data := cus.readfile("./dummy_data1.csv")
	fmt.Println(data[1][2])
}

type Customer struct {
	userID        string
	orderID       int
	FirstName     string
	LastName      string
	UserName      string
	Email         string
	UserType      int
	OrderAmount   []float64
	OrderDiscount []float64
}

func (r Customer) readfile(filepath string) [][]string {
	csvFile, err := os.Open(filepath)
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(csvFile)
	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return csvData
}
Marshal