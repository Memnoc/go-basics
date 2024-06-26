package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	fmt.Println("== Interfaces Assignment ==")

	tr := triangle{base: 10, height: 10}
	sq := square{sideLength: 10}

	printArea(tr)
	printArea(sq)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
