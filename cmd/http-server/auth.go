package main

import (
	"fmt"
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
		sendErrorResponse(rw, http.StatusBadRequest, nil, "inputs creds are not matching")
		return
	}
	sendResponse(rw, http.StatusOK, token, "Logged in succesfully")

}

func (app *App) handleLogout(rw http.ResponseWriter, r *http.Request) {
	// pick token from context that is set by middleware
	// delete it from redis
}

func (app *App) checkAllowance(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		token, err := getCookie(r, "token")
		if err != nil {
			//No cookie found
			sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized")
			return
		}
		categories := app.authService.CheckSession(r.Context(), token)
		fmt.Println(categories)
	})
}
