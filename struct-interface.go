package main

import (
	"fmt"
	"time"
)

// Interface, Composition, Method Overriding

func main() {
	// Gamze, Batuhan, Ayberk, Busra
	gamze := Developer{
		Employee: Employee{
			FirstName: "Gamze",
			LastName:  "Karadayı",
			Dob:       time.Date(2000, time.February, 15, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Senior Go Developer",
			Location:  "İzmir",
		},
		Skills: []string{"Go", "Docker", "Kubernetes", "C#"},
	}
	ellerhavada_batuhan := Developer{
		Employee: Employee{
			FirstName: "Batuhan",
			LastName:  "D.",
			Dob:       time.Date(1999, time.February, 18, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Junior Go Developer",
			Location:  "İzmir",
		},
		Skills: []string{"Go", "MySQL", "PHP", "ANV"},
	}
	ayberk := Developer{
		Employee: Employee{
			FirstName: "Ay",
			LastName:  "Bi",
			Dob:       time.Date(1998, time.July, 14, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Go Rapper",
			Location:  "İstanbul",
		},
		Skills: []string{"Go", "Python", "PostgreSQL", "Rust"},
	}
	mslogic := Manager{
		Employee: Employee{
			FirstName: "Bayan",
			LastName:  "Mantik",
			Dob:       time.Date(1987, time.March, 13, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Team Manager",
			Location:  "Yozgat",
		},
		Projects:  []string{"Face Detection", "e-Commerce"},
		Locations: []string{"San Francisco", "Yozgat"},
	}

	team := Team{
		"Go",
		"Golang Engineering Team",
		[]TeamMember{gamze, ayberk, ellerhavada_batuhan, mslogic},
	}
	team.PrintTeamDetails()
}

// Interface

type TeamMember interface {
	PrintName()
	PrintDetails()
}

// Struct -> Employee

type Employee struct {
	FirstName, LastName string
	Dob                 time.Time
	JobTitle, Location  string
}

func (e Employee) PrintName() {
	fmt.Printf("\n%s %s\n", e.FirstName, e.LastName)
}

func (e Employee) PrintDetails() {
	fmt.Printf("Date of Birth: %s, Job: %s, Location: %s\n", e.Dob.String(), e.JobTitle, e.Location)
}

// Struct -> Developer

type Developer struct {
	Employee // type embedding for composition
	Skills   []string
}

func (d Developer) PrintDetails() {
	d.Employee.PrintDetails()
	// d.PrintDetails()
	fmt.Println("Technical Skills: ")
	for _, v := range d.Skills {
		fmt.Println("-> " + v)
	}
}

// Struct -> Manager

type Manager struct {
	Employee
	Projects  []string
	Locations []string
}

func (m Manager) PrintDetails() {
	m.Employee.PrintDetails()
	fmt.Println("Projects: ")
	for _, v := range m.Projects {
		fmt.Println(v)
	}
	fmt.Println("Managing teams for the locations: ")
	for _, v := range m.Locations {
		fmt.Println(v)
	}
}

// Struct -> Team

type Team struct {
	Name, Description string
	TeamMembers       []TeamMember
}

func (t Team) PrintTeamDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members: ")
	for _, v := range t.TeamMembers {
		v.PrintName()
		v.PrintDetails()
	}
}
