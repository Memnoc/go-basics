package main

import (
	"fmt"
)

type shape interface {
	printArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) printArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func (t triangle) printArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func printArea(s shape) {
	fmt.Println(s.printArea())
}

func main() {
	fmt.Println("== Interfaces Assignment ==")

	tr := triangle{base: 4, height: 66}
	sq := square{sideLength: 4}

	printArea(tr)
	printArea(sq)
}
