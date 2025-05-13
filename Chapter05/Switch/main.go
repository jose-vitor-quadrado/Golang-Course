package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number from 1 to 3: ")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("Number 1")
	case 2:
		fmt.Println("Number 2")
	case 3:
		fmt.Println("Number 3")
	default:
		fmt.Println("Invalid number")
	}
}
