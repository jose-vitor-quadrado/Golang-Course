package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func main() {
	var rec rect
	rec.width = 5
	rec.height = 10

	r := rect{
		width:  20,
		height: 40,
	}

	fmt.Println(rect.area(rec))
	fmt.Println(rect.area(r))
}
