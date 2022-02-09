package main

import (
	"fmt"
	"math"
)

//  Define interfaces
type Shape interface {
	// interface method for structs, name area() that return float64
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// area method for Circle struct
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// area metho for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 7}
	rectangle := Rectangle{width: 7, height: 7}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
	fmt.Println(circle.area())
}
