package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i + 1)
	// }
	sum := func(x, y int) int {
		return x + y
	}
	fmt.Println(sum(5, 7))
}
