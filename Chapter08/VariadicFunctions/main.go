package main

import "fmt"

func sum(nums ...int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num += nums[i]
	}
	return num
}

func main() {
	fmt.Println(sum(10, 5, 10, 5, 15))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(sum(8))
}
