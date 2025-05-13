package main

import (
	"fmt"
	"math"
)

// Implements automatically
type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2*r.height + 2*r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func main() {
	rectangle := rect{
		height: 7,
		width:  9,
	}

	circle := circle{
		radius: 5,
	}

	fmt.Println(rectangle.area())
	fmt.Println(rectangle.perimeter())
	fmt.Println(circle.area())
	fmt.Println(circle.perimeter())
}
