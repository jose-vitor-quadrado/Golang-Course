package main

import "fmt"

func main() {
	var plantHeight int = 1
	for plantHeight < 5 {
		fmt.Println("still growing! current height:", plantHeight)
		plantHeight++
	}
	fmt.Println("plant has grown to", plantHeight, "inches")
}
