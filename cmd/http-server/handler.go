package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handler(r *chi.Mux) {
	r.Get("/web", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(rw, r, "./build/index.html")
	})
	r.Handle("/static/*", http.FileServer(http.Dir("./build")))
}
