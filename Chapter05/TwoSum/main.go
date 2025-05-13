package main

import "fmt"

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	indexes := make([]int, 2)

	for key, value := range nums {
		diff := target - value

		if index, ok := table[diff]; ok {
			indexes[0] = index
			indexes[1] = key
			break
		}

		table[value] = key
	}

	return indexes
}

func main() {
	var target int = 9
	nums := []int{11, 15, 2, 7}

	fmt.Println(twoSum(nums, target))
}
