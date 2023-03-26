package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	authservice "github.com/yards22/lcmanager/internal/auth_service"
)

type UserDetails struct {
	Token    string `json:"token"`
	Polls    bool
	Feedback bool
	Stories  bool
	MailId   string
}

type LoginRes struct {
	Token  string `json:"token"`
	MailId string `json:"mail_id"`
}

type Stories struct{}

type Token struct{}

type MailId struct{}

type Feedback struct{}

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

func (app *App) handleAddRole(rw http.ResponseWriter, r *http.Request) {
	var incBody authservice.RegisterRoleArgs
	err := getBody(r, &incBody)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}
	if incBody.MailId == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "mailid_missing")
		return
	}
	if incBody.SecretKey == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "secret_missing")
		return
	}
	if incBody.Role == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "role_missing")
		return
	}

	response := app.authService.PerformRoleSignup(r.Context(), incBody)

	if response != uuid.Nil.String() {
		sendResponse(rw, http.StatusOK, response, response)
		return
	}
	sendErrorResponse(rw, http.StatusUnauthorized, nil, "secret_invalid")
}

func (app *App) handleMe(rw http.ResponseWriter, r *http.Request) {
	var incBody string
	err := getBody(r, &incBody)
	if err != nil {
		fmt.Println(err)
	}
	params := app.authService.CheckSession(r.Context(), incBody)
	sendResponse(rw, http.StatusOK, params.MailID, "Check ME")
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
	var res LoginRes
	res.Token = app.authService.PerformLogin(r.Context(), incBody)
	if res.Token == uuid.Nil.String() {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "inputs creds are not matching")
		return
	}
	res.MailId = incBody.MailId
	sendResponse(rw, http.StatusOK, res, "Logged in succesfully")

}

func (app *App) handleLogout(rw http.ResponseWriter, r *http.Request) {

	token := (r.Context().Value("user")).(UserDetails).Token
	fmt.Println(token)
	app.authService.PerformLogout(r.Context(), token)
	sendResponse(rw, http.StatusOK, nil, "Logged out successfully")
}

func (app *App) checkAllowance(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := BearerAuthHeader(authHeader)
		fmt.Println(token)
		if token != "" {
			params := app.authService.CheckSession(r.Context(), token)
			fmt.Println(params)
			if params.OpenTo != nil {
				var newCtx context.Context
				var x UserDetails
				x.MailId = params.MailID
				x.Token = token
				for i := 0; i < len(params.OpenTo); i++ {
					if params.OpenTo[i] == "polls" {
						x.Polls = true
					}
					if params.OpenTo[i] == "stories" {
						x.Stories = true
					}
					if params.OpenTo[i] == "feedback" {
						x.Feedback = true
					}
				}
				newCtx = context.WithValue(r.Context(), "user", x)
				fmt.Println(newCtx.Value("user"))
				next.ServeHTTP(rw, r.WithContext(newCtx))
				return
			}
		}
		fmt.Println("here at unauthorized")
		sendErrorResponse(rw, http.StatusUnauthorized, nil, "unauthorized_user")
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
