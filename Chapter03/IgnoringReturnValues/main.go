package main

import "fmt"

func getNames() (string, string) {
	return "John", "Doe"
}

func main() {
	// get an error because we cannot compile with unused variables so change the name to underscore "_"
	// firstName, lastName := getNames()
	firstName, _ := getNames()
	fmt.Println("Welcome to Textio,", firstName)
}
