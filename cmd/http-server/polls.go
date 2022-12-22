package main

import (
	"net/http"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
)

func (app *App) handleCreatePoll(rw http.ResponseWriter, r *http.Request) {
	var arg sqlc.CreatePollsParams
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
}

func (app *App) handleGetPoll(rw http.ResponseWriter, r *http.Request) {

	entry := r.Context().Value(Pools{})

	if entry == true {

		polls := app.PollManager.Get(r.Context())
		sendResponse(rw, http.StatusCreated, polls, "These are the polls")
		return
	}

	sendErrorResponse(rw, http.StatusUnauthorized, nil, "Unauthorized for this route")

}
