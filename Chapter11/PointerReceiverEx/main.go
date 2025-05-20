package main

import "fmt"

type circle struct {
	x      int
	y      int
	radius int
}

func (c *circle) grow() {
	c.radius *= 2
}

func main() {
	c := circle{
		x:      1,
		y:      2,
		radius: 4,
	}

	c.grow()
	fmt.Println(c.radius)
}
