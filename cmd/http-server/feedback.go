package main

import (
	"fmt"
	"net/http"

	"github.com/yards22/lcmanager/internal/feedback_manager"
)

func (app *App) handleGetFeedback(rw http.ResponseWriter, r *http.Request) {
	x := (r.Context().Value("user")).(UserDetails)
	if x.Feedback {
		feedback := app.FeedbackManager.GetFeedback(r.Context())
		sendResponse(rw, http.StatusCreated, feedback, "feedback_section")
		return
	}

	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")
}

func (app *App) handlePostCommentFeedback(rw http.ResponseWriter, r *http.Request) {

	x := (r.Context().Value("user")).(UserDetails)

	if x.Feedback {
		var incBody feedback_manager.UpdateCommentsParams
		err := getBody(r, &incBody)

		if err != nil {
			fmt.Println(err)
		}
		app.FeedbackManager.PostCommentFeedback(r.Context(), incBody)
		sendResponse(rw, http.StatusCreated, nil, "comment_updated")
		return

	}
	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")
}
