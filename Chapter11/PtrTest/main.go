package main

import "fmt"

func main() {
	var num int = 5
	var ptr *int = &num

	*ptr = 10

	fmt.Println(num)
	fmt.Println(*ptr)
}
