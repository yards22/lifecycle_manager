package main

import (
	"net/http"

	"github.com/google/uuid"
	authservice "github.com/yards22/lcmanager/internal/auth_service"
)

func (app *App) handleSendOTP(rw http.ResponseWriter, r *http.Request) {
	var incBody authservice.SendOTPArgs
	err := getBody(r, &incBody)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}
	if incBody.MailId == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "mailid missing")
		return
	}

	response := app.authService.PerformMailIdCheck(r.Context(), incBody)

	if response != uuid.Nil.String() {
		sendResponse(rw, http.StatusOK, response, "OTP Sent Successfully")
		return
	}
	sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized")
	return
}

func (app *App) handleLogin(rw http.ResponseWriter, r *http.Request) {
	var incBody authservice.LoginArgs
	err := getBody(r, &incBody)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}
	if incBody.MailId == "" || incBody.OTP == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "inputs creds missing")
		return
	}

	token := app.authService.PerformLogin(r.Context(), incBody)

	if token == uuid.Nil.String() {

	}

}
