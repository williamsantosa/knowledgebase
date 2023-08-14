package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hi, my name is Will and this file was created using Go!")
	file.Close()

	stream, err := os.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)

	fmt.Println(s1)

	os.Remove("sample.txt")
}
