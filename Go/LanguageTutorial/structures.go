package main

import "fmt"

func main() {
	rect1 := Rectangle{30, 40}
	fmt.Printf("%v %v\n", rect1.height, rect1.width)
	fmt.Println(rect1.area())
}

// Structure like in C++
type Rectangle struct {
	height float64
	width  float64
}

// Function to return the width and height
// Takes input rectangle and is packaged as function area
// Similar to Rectangle::area() in C++
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}
