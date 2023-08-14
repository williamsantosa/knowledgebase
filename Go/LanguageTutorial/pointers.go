package main

import "fmt"

// same pointer logic as C/C++, with * and &
// * is a pointer to something (contains the memory address of smth)
// & is the memory address

func main() {
	x := 10

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", &x)

	changeValue(&x)

	fmt.Printf("%v\n", x)

}

// the type is always declared after the variable name
// Dereference the pointer by prefixing with * like C/C++
func changeValue(x *int) {
	*x = 7
}
