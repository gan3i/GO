package main

import "fmt"

type shapes interface {
	area() float64
}
type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		height : 5,
		base : 4,
	}
	printArea(t)

	s := square{
		sideLength : 10,
	}

	printArea(s)
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shapes) {
	fmt.Println(s.area())
}
