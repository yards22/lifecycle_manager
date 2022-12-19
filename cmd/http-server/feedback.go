package main

import (
	"net/http"
)

func (app *App) handleGetFeedback(rw http.ResponseWriter, r *http.Request) {

	data := app.FeedbackManager.GetFeedback(r.Context())

	sendResponse(rw, http.StatusCreated, data, "Feedbacks")
}
