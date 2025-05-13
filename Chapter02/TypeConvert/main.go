package main

import "fmt"

func main() {
	temperatureInt := 88
	temperatureFloat := float64(temperatureInt)

	fmt.Printf("%T - %T\n", temperatureInt, temperatureFloat)
}
