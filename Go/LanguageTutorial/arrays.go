package main

import "fmt"

func main() {
	// var EvenNum [5]int{0,2,4,6,8}

	// var EvenNum [5]int // declare size and type
	// for i := 0; i < len(EvenNum); i++ {
	// 	EvenNum[i] = i * 2
	// }

	// declare even num with that
	// range is exactly the same as enumerate in python
	// (returns 2 values for each iteration, first index and second copy of element)
	var EvenNum = [5]int{0, 2, 4, 6, 8}
	for i, value := range EvenNum {
		fmt.Println(value, i)
	}

	// notation to not use the i
	// no need to declare size
	var EvenNum2 = []int{0, 2, 4, 6, 8}
	for _, value := range EvenNum2 {
		fmt.Println(value)
	}

	// slicing array
	// same as python
	numSlice := []int{0, 2, 4, 6, 8}
	sliced := numSlice[0:3] // reference to it
	fmt.Printf("%v\n", sliced)

	// Make creates a new slice (dynamically sized array)
	// Capacity of new slice is the third argument to make
	slice2 := make([]int, 5, 10) // new slice, first argument defines datatype

	// Copies from source to destination (dest, source)
	copy(slice2, numSlice)
	fmt.Printf("%v\n%d\n", slice2, numSlice)

	// append to end
	slice3 := append(numSlice, 3, 2)
	fmt.Printf("%v\n", slice3)

}
