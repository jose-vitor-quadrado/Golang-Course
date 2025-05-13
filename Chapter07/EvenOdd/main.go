package main

import "fmt"

func main() {
	z := 0
	x := 8
	y := 9
	fmt.Scan(&z)

	if z%2 == 0 {
		par := SimpleMultiplication1(z, x)
		fmt.Println(par)
	}
	if z%2 != 0 {
		impar := SimpleMultiplication2(z, y)
		fmt.Println(impar)
	}
}

func SimpleMultiplication1(z, x int) int {
	return z * x
}

func SimpleMultiplication2(z, y int) int {
	return z * y
}
