package main

import (
	"net/http"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
)

func (app *App) handleCreatePoll(rw http.ResponseWriter, r *http.Request) {
	x := (r.Context().Value("user")).(UserDetails)
	if x.Polls {
		var arg sqlc.CreatePollsParams
		arg.PollBy = x.MailId
		err := getBody(r, &arg)
		if err != nil {
			sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
			return
		}

		app.PollManager.Create(r.Context(), arg)

		if err != nil {
			sendErrorResponse(rw, http.StatusInternalServerError, err, "")
			return
		}
		sendResponse(rw, http.StatusCreated, arg, "poll created")
		return
	}

	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")
}

func (app *App) handleGetPoll(rw http.ResponseWriter, r *http.Request) {

	x := (r.Context().Value("user")).(UserDetails)

	if x.Polls {

		polls := app.PollManager.Get(r.Context())
		sendResponse(rw, http.StatusCreated, polls, "polls_section")
		return
	}

	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")

}

