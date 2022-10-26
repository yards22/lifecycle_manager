package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handler(app *App, r *chi.Mux) {
	r.Get("/web", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(rw, r, "./build/index.html")
	})
	r.Handle("/static/*", http.FileServer(http.Dir("./build")))

	// Manager handlers
	r.Post("/{name}/close", func(rw http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			sendErrorResponse(rw, http.StatusBadRequest, nil, "Missing: manager name")
			return
		}
		app.managers[name].Close()
		sendResponse(rw, http.StatusOK, nil, "Success: Close "+name)
	})

	r.Post("/{name}/start", func(rw http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			sendErrorResponse(rw, http.StatusBadRequest, nil, "Missing: manager name")
			return
		}
		go app.managers[name].Run()
		sendResponse(rw, http.StatusOK, nil, "Success: Run "+name)
	})
}
