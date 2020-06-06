package main

import (
	"fmt"
	"strings"
)

// Declare constant
const Title = "Person Details"

// Declare package variable
var Country = "Out Of This World !!"

func main() {
	name := "Ankit Gaur"
	splitname := strings.Split(name, string(' '))
	fname := splitname[0]
	lname := splitname[1]
	age := 14
	// Print constant variable
	fmt.Println(Title)
	// Print local variables
	fmt.Println("First Name:", fname)
	fmt.Println("Last Name:", lname)
	fmt.Println("Age:", age)
	// Print package variable
	fmt.Println("Country:", Country)
}
