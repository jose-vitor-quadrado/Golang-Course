package main

import "fmt"

func naive(arr []int, size int) {
	temp := make([]int, size)
	for i := 0; i < size; i++ {
		temp[i] = arr[size-i-1]
	}
	for i := 0; i < size; i++ {
		arr[i] = temp[i]
	}
	fmt.Println(arr)
}

func reverseGo(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)
}

func swapPtrs(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func pointers(arr []int, size int) {
	left, right := 0, size-1
	for left < right {
		swapPtrs(&arr[left], &arr[right])
		left++
		right--
	}
	fmt.Println(arr)
}

func swapElements(arr []int, size int) {
	for i := 0; i < size/2; i++ {
		temp := arr[i]
		arr[i] = arr[size-i-1]
		arr[size-i-1] = temp
	}
	fmt.Println(arr)
}

func recursion(arr []int, left int, right int) {
	if left >= right {
		return
	}
	temp := arr[left]
	arr[left] = arr[right]
	arr[right] = temp
	recursion(arr, left+1, right-1)
}
func reverseRec(arr []int, size int) {
	recursion(arr, 0, size-1)
	fmt.Println(arr)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	var size int = len(arr)

	naive(arr, size)
	reverseGo(arr)
	pointers(arr, size)
	swapElements(arr, size)
	reverseRec(arr, size)
}
