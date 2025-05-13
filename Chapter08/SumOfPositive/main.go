package main

import "fmt"

func main() {
	arr := []int{1, -4, 7, 12}
	sum := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			continue
		}
		sum += arr[i]
	}
	fmt.Println(sum)
}
