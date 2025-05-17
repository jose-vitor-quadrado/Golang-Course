package main

import "fmt"

func sumSlice(lista []int) int {
	sum := 0
	for _, num := range lista {
		sum += num
	}
	return sum
}

func main() {
	var n int
	fmt.Print("quantos numeros a lista tera? ")
	fmt.Scan(&n)
	lista := []int{}
	for i := 0; i < n; i++ {
		lista = append(lista, i)
	}
	pares := []int{}

	for i := 0; i < len(lista); i++ {
		if i%2 != 0 || i == 0 {
			continue
		}
		pares = append(pares, lista[i])
	}

	fmt.Println(pares)
	fmt.Println(sumSlice(pares))
}
