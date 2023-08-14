package main

import "fmt"

// Defer statement defers execution until surrounding functions return
// Multiple defers are pushed into stack in LIFO order
// Defer used to cleanup resources like file, database connections, etc

// Recover helps regain normal flow of execution after panic
// Used with defer to reecover panic in goroutine

// Panic throws exception in other languages
// Normal execution flows stops immediately
// Deferred function are executed normally

func main() {
	// First then second, normal
	FirstRun()
	SecondRun()

	Separator()

	// Defer causes SecondRun (surrounding func) to run first
	defer FirstRun()
	SecondRun()

	// Recovery
	fmt.Println(Div(3, 0)) // errors because of division by 0
	fmt.Println(Div(3, 5))
	demPanic()

	// Exits using approviate status because not in a function
	panic("Panic 2")
}

func FirstRun()  { fmt.Println("Execute FirstRun") }
func SecondRun() { fmt.Println("Execute SecondRun") }
func Separator() { fmt.Println("--------------------") }

func Div(num1 int, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

// Panic within function does not break program
func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}
