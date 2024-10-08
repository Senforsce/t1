package main

import (
	"fmt"
	"net/http"

	"github.com/senforsce/tndr"
)

func main() {
	component := hello("John")

	http.Handle("/", t1.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
