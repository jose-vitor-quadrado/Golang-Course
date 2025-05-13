package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "grape"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}
