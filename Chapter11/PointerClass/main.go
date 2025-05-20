package main

import "fmt"

func main() {
	var x int = 5
	var p *int = &x

	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(x)
	*p = 10
	fmt.Println(*p)
	fmt.Println(p)
	fmt.Println(x)
}
