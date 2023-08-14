# Goroutines

Lightweight thread managed by Go runtime. `go f(x, y, z)` starts a new goroutine running `f(x, y, z)`.

```Go
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

```

Channels are a typed conduit which can send and receive values with the channel operator, `<-`. 

```Go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

Theyt must be created before use.

```Go
ch := make(chan int)
```

```Go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
```

Channels can be buffered. Provide the buffer length as the second argument to make a buffered channel. Buffered channels block only when buffer is full. `ch := make(chan int, 100)`.

A sender can `close` a channel to indicate that no more values will be sent. Can be tested with `v, ok := <-ch`.

The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run and then executes the case. One is randomly chosen if multiple are ready. The default is selected if no other case is ready.

```Go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

Mutex with standard library in `sync.Mutex`. Has `Lock` and `Unlock`.