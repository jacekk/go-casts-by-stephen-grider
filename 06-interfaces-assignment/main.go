package main

import (
	"fmt"
)

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	fmt.Print("Square: ")
	printArea(square{sideLength: 2})
	fmt.Print("Triangle: ")
	printArea(triangle{base: 3, height: 2})
}
