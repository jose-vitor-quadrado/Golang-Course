package main

import "fmt"

func reverseArr(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func binaryCalc(num int) []int {
	arr := []int{}
	i := 0
	for num > 0 {
		arr = append(arr, num%2)
		num /= 2
		i++
	}
	return reverseArr(arr)
}

func arrToStringConverter(arr []int) string {
	var binary string = ""
	for i := 0; i < len(arr); i++ {
		binary += fmt.Sprintf("%d", arr[i])
	}
	return binary
}

func main() {
	number, binArr := 45, []int{}
	var result string
	if number < 0 {
		fmt.Println("Invalid number")
		return
	}
	binArr = binaryCalc(number)
	result = arrToStringConverter(binArr)
	fmt.Println(result)
}
