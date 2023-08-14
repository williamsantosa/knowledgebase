package main

import "fmt"

// Numeric
// - int
// - float
// String
// Boolean
// - true
// - false
// Derived
// - pointer
// - array
// - structure
// - map
// - interface

func main() {

	// const/var name type = value
	const pi float64 = 3.1415

	// if no value is assigned, must declare the type
	var varNoValue int // declare
	varNoValue = 10    // assign
	// if value assigned, does not need to declare type
	var varValue = 5 // lefthand side (var varValue) declare, = 5 assign
	// can declare this way as well (declaration + assignment)
	// it will infer the type of shorthandValue
	shorthandValue := 10

	// Multiple variable definition
	var (
		varA = 2
		varB = 3
	)

	// Strings use " or ` to delimit them

	var Name string = "William Santosa"

	fmt.Println(len(Name)) // print out the length
	fmt.Printf("%v %v %v %v %v %v", varNoValue, varValue, varA, varB, Name, shorthandValue)

}
