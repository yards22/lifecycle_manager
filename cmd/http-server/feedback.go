package main

import (
	"fmt"
	"net/http"
)

func (app *App) handleGetFeedback(rw http.ResponseWriter, r *http.Request) {

	entry := r.Context().Value(Feedback{})

	fmt.Println("entry_bool", entry)

	if entry == true {

		feedback := app.FeedbackManager.GetFeedback(r.Context())
		sendResponse(rw, http.StatusCreated, feedback, "feedback_section")
		return
	}

	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")
}
