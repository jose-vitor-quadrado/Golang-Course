package main

import "fmt"

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

func main() {
	names := []string{"Logan", "Sue Storm", "Tony Stark"}
	printStrings(names...)
}
