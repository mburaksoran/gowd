package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := Person{
		ID:        7,
		FirstName: "mehemed",
		LastName:  "soran",
		UserName:  "memothelegend",
		Gender:    "Erkek",
		Name:      Name{Family: "Sorkanli", Personal: "memo"},
		Email: []Email{
			{ID: 1, Kind: "work", Address: "memo@mail.com"},
			{ID: 2, Kind: "home", Address: "memom@mail.com"},
		},
		Interest: []Interest{
			{ID: 1, Name: "c#"},
			{ID: 2, Name: "asd"},
			{ID: 3, Name: "asddda"},
			{ID: 4, Name: "ssss"},
			{ID: 5, Name: "dlasdl"},
		},
	}
	WriteMessage("personal fullname")
	WriteStarLine()
	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()
	WriteMessage("\n")
	WriteMessage("personal email with index")
	WriteStarLine()
	resEmail := GetPersonEmailAddress(&person, 0)
	WriteMessage(resEmail)
	WriteMessage("\n")
	WriteMessage("Personal Email Object with Index")
	WriteStarLine()
	resEmail2 := GetPersonEmail(&person, 0)
	fmt.Println(resEmail2)
	WriteStarLine()
	WriteMessage("Reading Operation Ended.")
	WriteMessage("\n")
	WriteMessage("Writing Operation Started.")
	SaveJSON("person.json", person)
	WriteMessage("writing operation ended.")

}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

type Name struct {
	Family   string
	Personal string
}
type Email struct {
	ID      int
	Kind    string
	Address string
}
type Interest struct {
	ID   int
	Name string
}

func GetPerson(p *Person) string {
	return p.FirstName + "  " + p.LastName
}

func GetPersonEmailAddress(p *Person, ind int) string {
	return p.Email[ind].Address
}

func GetPersonEmail(p *Person, ind int) Email {
	return p.Email[ind]
}
func WriteMessage(msg string) {
	fmt.Println(msg)
}
func WriteStarLine() {
	fmt.Println("****************")
}
func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}
