package main

import "fmt"

func yearsUntilEvents(age int) (
	yearsUntilAdult int,
	yearsUntilDrinking int,
	yearsUntilCarRental int,
) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}

func printMessage(age int) {
	isAdult, isAbleToDrink, isAbleToRentACar := yearsUntilEvents(age)

	fmt.Println("Age:", age)
	fmt.Printf("You are and adult in %d years\n", isAdult)
	fmt.Printf("You can drink in %d years\n", isAbleToDrink)
	fmt.Printf("You can rant a car in %d years\n", isAbleToRentACar)
	fmt.Println("===========================================")
}

func main() {
	printMessage(10)
	printMessage(19)
	printMessage(22)
	printMessage(30)
}
