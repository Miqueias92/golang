package main

import (
	"fmt"
)

var name string

func main() {
	name = "Miqueias"
	var middleName = "Santos"
	lastName := "Sousa"
	year := 2024

	// first way
	fmt.Println("Hello", name+" "+middleName+" "+lastName)

	// second way
	fmt.Println("Hello", name, middleName, lastName)

	// third way
	fmt.Printf("Welcome to %d \n", year)
}
