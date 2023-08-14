package main

import "fmt"

// Uses the same operators as C/C++
// https://go.dev/ref/spec#Operators_and_punctuation

func main() {
	fmt.Printf("%v %v", false && true, true || false)
}
