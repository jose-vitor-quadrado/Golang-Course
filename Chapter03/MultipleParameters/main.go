package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func message() {
	fmt.Println("Salve")
}

func main() {
	fmt.Println(add(10, 5))
	fmt.Println(sub(10, 5))
	message()
}
