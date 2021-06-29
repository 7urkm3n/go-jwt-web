package main

import (
	"net/http"
	"time"

	"github.com/pascaldekloe/jwt"
)

func (app *application) profileHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	err := app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) signInHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user, err := app.models.Users.GetByEmailWithPassword(input.Email)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	// validate password
	if user.Password != input.Password {
		app.failedValidationResponse(w, r, map[string]string{"errors": "email or password not matched"})
		return
	}

	// Token Generate
	var claims jwt.Claims
	claims.Subject = input.Email
	claims.Issued = jwt.NewNumericTime(time.Now())
	claims.NotBefore = jwt.NewNumericTime(time.Now())
	claims.Expires = jwt.NewNumericTime(time.Now().Add(24 * time.Hour))
	claims.Issuer = "sample"
	claims.Audiences = []string{"sample"}

	jwtBytes, err := claims.HMACSign(jwt.HS256, []byte(app.config.jwt.secret))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Convert the []byte slice to a string and return it in a JSON response.
	err = app.writeJSON(w, http.StatusCreated, envelope{"payload": string(jwtBytes)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
