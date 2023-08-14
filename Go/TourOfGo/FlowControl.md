# Flow Control

The basic `for`` loop has three components separated by semicolons:
1. the init statement: executed before the first iteration
2. the condition expression: evaluated before every iteration. Optional.
3. the post statement: executed at the end of every iteration. Optional.

```Go
for i := 5; i < 20; i++ {
  print(i)
}
```

While loop.

```Go
i := 5
for i < 20 {
  print(i)
  i++
}
```

Infinite loop.

```Go
for {
  print("hi")
}
```

If statement.

```Go
if i == true {
  print("true!")
} else if i == false{
  print("false!")
} else {
  print("else!")
}
```

The if statement can begin with a short statement to execute before condition.

```Go
if v := n * 2; v < n + 5 {
  return v;
}
```

Variables declared inside of the `if` short statement are available in the other `else`/`elseif` blocks.

Switch statment. Top to bottom. Stops when case succeeds, unlike C/C++.

```Go
switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
```

Switch without condition is the same as `switch true`. Can use to clean long if then else chains.

```Go
t := time.Now()
switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
```

Defer statement defers execution of function until surrounding function returns. Pushed onto a stack, LIFO.