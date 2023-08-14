# Basics

Go names are exported if it begins ith a capital letter.

`x int, y int` is the same as `x, y int`.

You can return multiple results in Go. For example,
```Go
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

A return statement without arguments returns the named return values. This is known as a "naked" return.

```Go
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
```

Can initialize variables multiple.

```Go
var i, j int = 1, 2
```

Basic types.

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Variables declare without explicit inital value are given the zero value.

`0` for numeric types,
`false` for the boolean type, and
`""` (the empty string) for strings.

The express `T(v)` convert value `v` to type `T`, for example,
```Go
var x int = 5
var y float64 = float64(x)
```
converts x into float64 for variable y.
Assignment between items of different type requires an explicit conversion. In C, you can assign them and then it'll infer it as whatever you're inferring it as in its stored format type.

Declaring a variable without specifying an explicit type allows the type to be inferred from the value on the righthand side.

Numeric constants are high-precision values.

