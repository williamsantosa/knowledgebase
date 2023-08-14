# Pointers

Pointer.

```Go
var p *int
```

The & operator generates a pointer (stores mem address) to its operand.

The `*` operator denotes pointer's underlying value. Known as "dereferencing" or "indirecting".

There is no pointer arithmetic in Go ):

A `struct` is a collection fields. They are accessed using a dot. Struct fields can be accessed through a struct pointer.

To access the field `X` of a struct when we have the struct pointer p we could write `(*p).X`. However, that notation is cumbersome, so the language permits us instead to write just `p.X`, without the explicit dereference.

```Go
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
```

Struct literals denote a newly allocated struct value by listing the values of its fields. Can list a subset of fields by using `Name:` syntax.

```Go
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)
```

Array. `[n]T`. Sliced using `a[low:high]`. Excludes last but includes first. Slice dynamically sized. Does not store data, describes a section of underlying array. They reference the same array.

Slice literals are array literals without the length.

Array literal.
```Go
[3]bool{true, true, false}
```

Slice literal.

```Go
[]bool{true, true, false}
```

Slice has a length and capacity. The length and capacity can obtained using `len(s)` and `cap(s)`. Can extend a slice's length by re-slicing it if it has sufficient capacity.

Zero value of slice is `nil`.

Dynamically create a slice using `make` function. Allocates a zeroded array and returns a slice that refers to that array. 

```Go
a := make([]int, 5)  // len(a)=5
```

The capacity is the third argument. Optional.

```Go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

Slices can contain any type, including other slices.

Append. `func append(s []T, vs ...T) []T`. First parameter is slice of type `T`, rest are `T` values appended. Resulting value is slice containing elements of original plus provided. If the backing array of `s` is too small a new bigger array is allocated. returned slice points to newly allocated array.

The `range` form is the same as `enumerate` in Python. The first is index, second is element copy.

```Go
for i, _ := range pow     // only index
for _, value := range pow // only value
for i := range pow        // only index
```

Maps map keys to values. The zero value of a map is `nil`. No keys, nor keys can be added. A `make` function returns map of given type, initialized and ready for use. `var m = make(map[string]int)`. map with string as key and int as value.

Map literals are struct literals but keys are required.

```Go
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{ 40.68433, -74.39967, },
	"Google":    Vertex{ 37.42202, -122.08408, },
}
```

If top-level type is just the type name it can be omitted.

```Go
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": { 40.68433, -74.39967, },
	"Google":    { 37.42202, -122.08408, },
}
```

Insert or update an element in map m: `m[key] = elem`
Retrieve an element: `elem = m[key]`
Delete an element: `delete(m, key)`
Test that a key is present with a two-value assignment: `elem, ok = m[key]`

Functions can be used as arguments and return values.

Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

The `adder` function returns a closure. Each closure is bound to its own `sum` variable.

```Go
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
```

Interpret the above adder function as function with name `adder` that returns the type `func(int) int`. The function within `adder`, `func(x int) int`, references a variable outside of its scope defined at `sum := 0`. Therefore, the function value `sum` is a closure.

Another explanation. The function `adder` has a `sum` value that is used in an internal function. The `sum` value persists through calls since the `adder` function contains its own `sum` value that is unique from other `adder` functions called.

Fibonacci exercise.

```Go
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		temp := x + y
		x = y
		y = temp
		return y
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```