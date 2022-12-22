package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initHandler(app *App, r *chi.Mux) {
	r.Post("/sendOTP", app.handleSendOTP)
	r.Post("/login", app.handleLogin)
	r.Delete("/logout", app.checkAllowance(http.HandlerFunc(app.handleLogout)))
	r.Post("/poll", app.checkAllowance(http.HandlerFunc(app.handleCreatePoll)))
	r.Get("/poll", app.checkAllowance(http.HandlerFunc(app.handleGetPoll)))
	r.Get("/feedback", app.checkAllowance(http.HandlerFunc(app.handleGetFeedback)))
}
