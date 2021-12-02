package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	xmlFile, err := os.Open("Employees.xml")
	if err != nil {
		fmt.Println("opening file error : ", err)
		return
	}
	defer xmlFile.Close()
	xmlData, _ := ioutil.ReadAll(xmlFile)

	var c Company
	xml.Unmarshal(xmlData, &c) // xmlData ya gelen byte[] 'si formatındaki veriyi Company nin nesne örneğine aktarır.

	fmt.Println(c.Persons[:])
	fmt.Println(c.Persons[0] == c.Persons[1])

	var person jsonPerson
	var persons []jsonPerson

	for _, value := range c.Persons {
		person.ID = value.ID
		person.FirstName = value.FirstName
		person.LastName = value.LastName
		person.UserName = value.UserName
		persons = append(persons, person)
	}

	jsonData, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))

	jsonFile, err := os.Create("Employee.json")
	if err != nil {
		fmt.Println(err)
	}
	jsonFile.Write(jsonData)
	jsonFile.Close()
}

//xml to json convertion

type jsonPerson struct {
	ID        int    `xml:"id"`
	FirstName string `xml:"firstname"`
	LastName  string `xml:"lastname"`
	UserName  string `xml:"username"`
}

type Person struct {
	XMLNAME   xml.Name `xml:"person"`
	ID        int      `xml:"id"`
	FirstName string   `xml:"firstname"`
	LastName  string   `xml:"lastname"`
	UserName  string   `xml:"username"`
}

type Company struct {
	XMLNAME xml.Name `xml:"company"`
	Persons []Person `xml:"person"`
}

func (p Person) String() string {
	return fmt.Sprintf("\t ID: %d - FirstName : %s %s -> UserName : %s ", p.ID, p.FirstName, p.LastName, p.UserName)
}
