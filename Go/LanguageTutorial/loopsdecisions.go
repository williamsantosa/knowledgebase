package main

import "fmt"

func main() {
	// Loop with variable and step
	for i := 1; i < 5; i++ {
		fmt.Printf("%d", i)
	}

	// While loop
	i := 1
	for i <= 10 {
		fmt.Printf("%d", i)
		i += 1
	}

	// For-each loop
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}

	// Exiting loop
	// break and continue work the same as Python
	// Continue goes to next iteration
	// Break exits the current loop
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum) // 6 (2+4)

	// if else
	var age int = 18
	if age > 18 {
		print("age is over 18\n")
	} else if age >= 13 {
		print("age is between 13 and 18\n")
	} else {
		print("age is under 13\n")
	}

	// Switch statement

	switch age {
	case 18:
		print("age is 18\n")
	case 20:
		print("age is 20\n")
	default:
		print("age is not 20 or 18\n")
	}
}
