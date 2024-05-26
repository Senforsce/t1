package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", t1.Handler(Home()).ServeHTTP)
	http.ListenAndServe(":3000", r)
}
