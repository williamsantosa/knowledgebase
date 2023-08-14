package main

import "fmt"

func main() {
	num := 5
	fmt.Println(factorial(num))
}

// notice doesn't matter where the function gets placed, can still call in main
// no header file required either

// Return type after parentheses (if none, leave blank)
// Parameters within parentheses
func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1) // called the same way as other languages
}
