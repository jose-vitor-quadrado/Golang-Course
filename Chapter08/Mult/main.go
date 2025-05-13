package main

import "fmt"

func main() {
	var multTable [10]int
	var num int
	fmt.Scan(&num)
	for i := 0; i < 10; i++ {
		multTable[i] = num * (i + 1)
	}
	fmt.Println(multTable)
}
