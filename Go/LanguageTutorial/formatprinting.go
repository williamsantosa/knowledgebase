package main

import "fmt"

func main() {
	var name string = "William Santosa"
	const pi float64 = 3.1415
	win := true
	x := 5

	fmt.Printf("%.3f\n", pi)  // floating point
	fmt.Printf("%#v\n", name) // Go syntax representation of value
	fmt.Printf("%v\n", name)  // print value (inferred)
	fmt.Printf("%T\n", name)  // print value type
	fmt.Printf("%t\n", win)   // boolean
	fmt.Printf("%d\n", x)     // base 10 (decimal)
	fmt.Printf("%b\n", 25)    // binary
	fmt.Printf("%c\n", 34)    // char code
	fmt.Printf("%x\n", 15)    // hex code
}
