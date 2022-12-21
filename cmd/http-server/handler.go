package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initHandler(app *App, r *chi.Mux) {
	r.Post("/sendOTP", app.handleSendOTP)
	r.Post("/login", app.handleLogin)
	r.Post("/poll", http.HandlerFunc(app.handleCreatePoll))
	r.Get("/poll", http.HandlerFunc(app.handleGetPoll))
	r.Get("/feedback", http.HandlerFunc(app.handleGetFeedback))
}
