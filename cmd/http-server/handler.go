package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initHandler(app *App, r *chi.Mux) {
	r.Post("/sendOTP", app.handleSendOTP)
	r.Post("/login", app.handleLogin)
	r.Delete("/logout", app.checkAllowance(http.HandlerFunc(app.handleLogout)))
	r.Post("/addRole", app.handleAddRole)
	r.Post("/poll", app.checkAllowance(http.HandlerFunc(app.handleCreatePoll)))
	r.Get("/poll", app.checkAllowance(http.HandlerFunc(app.handleGetPoll)))
	r.Post("/stories/uploadImg", app.checkAllowance(http.HandlerFunc(app.handleUploadImage)))
	r.Get("/feedback", app.checkAllowance(http.HandlerFunc(app.handleGetFeedback)))
	r.Post("/feedback/comment", app.checkAllowance(http.HandlerFunc(app.handlePostCommentFeedback)))

	// WebUI
	r.Get("/web", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(rw, r, "./build/index.html")
	})
	r.Handle("/static/*", http.FileServer(http.Dir("./build")))
}
