# Methods

Method is a function with a special *receiver* argument.

The receiver appears in its own argument list between the `func` keyword and method name.

The `Abs` method has a receiver of type `Vertex` named `v`. The struct gets treated like a variable in the function and can be referenced. It is called like a class method in other languages. 

```Go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

A method is just a function with a receiver argument. The following is equivalent.

```Go
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

Methods can also be declared on non-structs, like `float64`. Only possible if receiver is defined in same package as meythod.

Pointer receivers are methods where the receiver type has the literal syntax `*T` for some type `T`. `T` cannot itself be a pointer. It modifies the value like a getter/setter. Without it, it would not.

Interfaces are a set of method signatures. They are implemented implicitly. There is no explicit declaration of intent and no "implements" keyword. It decouples the definition of an interface from its implementation, and can appear in anyt package without prearrangement.

Interface values can be thought of as a tuple of a value and a concrete type (`(value, type)`). An interface value holds a value of a specific underlying concrete type. Calling a method on an interface value executes the method of the same name on its underlying type. If `nil`, it is called with a nil receiver.

The empty interface may hold values of any type.

Type assertions provide access to interface value's underlying concrete value.

```Go
t := i.(T)
```

The interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`. T otest if it can hold a specific type, a type assertion can return two values: underlying value and a bool value reporting if it succeeded.

```Go
t, ok := i.(T)
```

Type switches are a construct that permits several type assertions in series.

```Go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

`Stringer` is an interface that can describe itself as a string. 

```Go
type Stringer interface {
  String() string
}
```

Exercise

```Go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0],ip[1],ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
```

Errors. `nil` means success, otherwise failure.

```Go
type error interface {
  Error() string
}
```

