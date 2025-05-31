package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func print_naming_conventions() {
	const MAXRETRIES = 3

	var employee1 Employee

	fmt.Println(employee1)
}
