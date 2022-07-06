package main

import (
	"fmt"
)

type employee struct {
	First string
	Last  string
	Age   int
}

type manager struct {
	employee
	First         string
}

func main() {

	mgmt:= manager{
		employee: employee{
			First:"Jordan", 
			Last: "Satterfield", 
			Age: 30,
		},
		First: "Manager Level 1",
	}
	// fields and methods of the inner-type are promoted to the outer-type
	fmt.Println(mgmt.First, mgmt.employee.First)
}
