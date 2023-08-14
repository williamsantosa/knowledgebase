# Generics

Go functions can work on multiple types using type parameters. The type parmeters appeara between brackets before the function's arguments.

```Go
func Index[T comparable](s []T, x T) int
```

This declaration means that s is a slice of any type T that fulfills the built-in constraint comparable. x is also a value of the same type. Useful since can use `==` and `!=` on same type.

A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.

```Go
package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
```