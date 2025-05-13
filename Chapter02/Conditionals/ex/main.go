package main

import "fmt"

func main() {
	var email string = ""

	/*
		length := len(email)
		if length < 1 {
			fmt.Println("Email is invalid")
		}
	*/
	// We can do
	if length := len(email); length < 1 {
		fmt.Println("Email is invalid")
	}
}
