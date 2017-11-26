package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type rectangle struct {
	sideLength float64
}

func main() {
	r := rectangle{
		sideLength: 2,
	}
	t := triangle{
		height: 12,
		base:   14,
	}
	printArea(r)
	printArea(t)
}

func printArea(s shape) {
	fmt.Printf("The Shape: %s has the Area: %v\n", reflect.TypeOf(s), s.getArea())
}

func (t triangle) getArea() float64 {
	return (0.5 * t.base * t.height)
}

func (r rectangle) getArea() float64 {
	return (r.sideLength * r.sideLength)
}