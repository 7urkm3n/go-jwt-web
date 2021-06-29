package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// router.HandlerFunc(http.MethodGet, "/v1/signIn", app.signInHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/tokens/authentication", app.authenticationHandler)
	router.HandlerFunc(http.MethodGet, "/v1/profile", app.requireAuth(app.profileHandler))
	router.HandlerFunc(http.MethodPost, "/v1/signin", app.signInHandler)

	return app.recoverPanic(app.authenticate(router))
}
