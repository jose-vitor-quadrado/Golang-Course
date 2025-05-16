package main

import "fmt"

func main() {
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 24
	ages["Mary"] = 21
	ages = map[string]int{
		"Wolverine": 280,
		"Magneto":   52,
	}

	fmt.Println(ages)
	fmt.Println(len(ages))
}
