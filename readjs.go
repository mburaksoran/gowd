package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main2() {
	csvFile, err := os.Open("./dummy_data1.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(csvFile)
	csvData, err := reader.ReadAll()
	fmt.Println(csvData[1][1])
	/* 	defer csvFile.Close()

	   	reader := csv.NewReader(csvFile)

	   	csvData, err := reader.ReadAll()
	   	if err != nil {
	   		fmt.Println(err)
	   		os.Exit(1)
	   	}

	   	var emp Employee
	   	var employees []Employee

	   	for _, each := range csvData {
	   		emp.FirstName = each[0]
	   		emp.LastName = each[1]
	   		emp.Age, _ = strconv.Atoi(each[2])
	   		employees = append(employees, emp)
	   	}

	   	// Print data
	   	// fmt.Println(employees)

	   	// Convert to JSON

	   	jsonData, err := json.Marshal(employees)
	   	if err != nil {
	   		fmt.Println(err)
	   		os.Exit(1)
	   	}
	   	fmt.Println(string(jsonData))

	   	// Create JSON

	   	jsonFile, err := os.Create("data.json")
	   	if err != nil {
	   		fmt.Println(err)
	   	}
	   	defer jsonFile.Close()
	   	jsonFile.Write(jsonData) */
}
