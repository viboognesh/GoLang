package basics

import "fmt"

var middleName = "Cane"

func print_variables() {
	// var age int
	// var name string = "John"
	// var name1 = "Jane"

	// count := 10
	// lastname := "Smith"

	fmt.Println(middleName)
	middleName := "Mayor"
	fmt.Println(middleName)
	printName()
	// default values

}

func printName() {
	firstName := "Michael"
	fmt.Println(firstName)
	fmt.Println(middleName) //TODO Hello There
}
