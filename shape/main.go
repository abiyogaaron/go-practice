package main

import "fmt"

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

func main() {
	tr := triangle{base: 10, height: 10}
	sq := square{sideLength: 10}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}
func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}
