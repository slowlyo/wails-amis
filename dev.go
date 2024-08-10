package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("frontend/dist/"))

	http.Handle("/", fs)

	fmt.Println("http://localhost:8080")

	_ = http.ListenAndServe(":8080", nil)
}
