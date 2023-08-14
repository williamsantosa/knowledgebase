package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{50, 60}
	circ := Circle{7}

	fmt.Println("Area of rectangle is", getArea(rect))
	fmt.Println("Area of rectangle is", getArea(circ))
}

// struct defines fields while interfaces define methods
// Provides method signature sfr similar types of objects
// Go does not have classes and inheritance so it uses this
type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
}

func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}

// Shape
// Either use it like this and pass it in as a parameter
// or use it by defining the type declared in interface using
// struct definition
func getArea(shape Shape) float64 {
	return shape.area()
}
