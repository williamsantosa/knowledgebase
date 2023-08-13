# Go Programming Language Tutorial

Following YouTube video located [here](https://youtu.be/Q0sKAMal4WQ).

## first.go

```Go
package main

import "fmt"

func main () {

  fmt.Println("Hello world!")

}
```

The `package main` tells the Go compiler that the package should compile as an executable program instead of a shared library. Therefore, the main function in package "main" is the entry point of the executable program. Building shared libraries does not have the main package or the main function.

The `fmt` import is the "format" (pronounched fumpt) that implements formatted I/O with functions similar to C's printf and scanf.