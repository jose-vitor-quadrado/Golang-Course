package main

import "fmt"

func main() {
	mySlice1 := make([]int, 5, 10)
	mySlice2 := make([]int, 5)
	mySlice3 := []string{"I", "love", "go"}

	fmt.Println(mySlice1)
	fmt.Println(mySlice2)
	fmt.Println(mySlice3)
	fmt.Println(len(mySlice1))
	fmt.Println(cap(mySlice1))
}
