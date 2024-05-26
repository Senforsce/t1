package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var flagPort = flag.Int("port", 0, "Set the HTTP listen port")

func main() {
	flag.Parse()

	if *flagPort == 0 {
		fmt.Println("missing port flag")
		os.Exit(1)
	}

	var count int
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		c := Page(count)
		t1.Handler(c).ServeHTTP(w, r)
	})
	err := http.ListenAndServe(fmt.Sprintf("localhost:%d", *flagPort), nil)
	if err != nil {
		fmt.Printf("Error listening: %v\n", err)
		os.Exit(1)
	}
}
