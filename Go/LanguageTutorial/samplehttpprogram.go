package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler2)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to my home page")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
