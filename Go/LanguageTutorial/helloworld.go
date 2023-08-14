// The `package main` tells the Go compiler that the package should compile as an executable program instead of a shared library.
// Therefore, the main function in package "main" is the entry point of the executable program.
// Building shared libraries does not have the main package or the main function.
package main

// The `fmt` import is the "format" (pronounched fumpt) that implements formatted I/O with functions similar to C's printf and scanf.
import "fmt"

// Main function that gets run
func main() {

	// Same as printf in C
	fmt.Println("Hello World! This is my first Go program.")

}
