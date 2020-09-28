package main

import (
	"fmt"
	"go-tutorials/internal/employee"
)

func main() {
	emp := employee.Employee{
		Name:    "Bob",
		Address: "New York, New york",
	}
	emp.SetID("1a2a3a4a5a6a")
	emp.Salary = 7000
	emp.Position = "Senior"
	emp.ManagerID = "1b2b3b4b5b6b"
	fmt.Printf("%+v\n", emp)
}
