package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	authservice "github.com/yards22/lcmanager/internal/auth_service"
)

type Pools struct{}

type Blogs struct{}

type Token struct{}

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
	token := r.Context().Value(Token{}).(string)
	app.authService.PerformLogout(r.Context(), token)
	// delete it from redis
	sendResponse(rw, http.StatusOK, nil, "Logged out successfully")
}

func (app *App) checkAllowance(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := BearerAuthHeader(authHeader)
		if token != "" {
			categories := app.authService.CheckSession(r.Context(), token)
			var newCtx context.Context
			newCtx = context.WithValue(r.Context(), Token{}, token)

			for i := 0; i < len(categories); i++ {
				if categories[i] == "polls" {
					newCtx = context.WithValue(r.Context(), Pools{}, true)
				}
				if categories[i] == "blogs" {
					newCtx = context.WithValue(r.Context(), Blogs{}, true)
				}
			}
			next.ServeHTTP(rw, r.WithContext(newCtx))
		}
		sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized user")
	})
}

func BearerAuthHeader(authHeader string) string {
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
