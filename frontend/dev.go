package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("dist/"))

	http.Handle("/", fs)

	fmt.Println("http://localhost:8090")

	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
