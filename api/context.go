package main

import (
	"context"
	"net/http"
	"sample/models/users"
)

// Define a custom contextKey type, with the underlying type string.
type contextKey string

const userContextKey = contextKey("user")

// contextSetUser: User struct added to the context.
func (app *application) contextSetUser(r *http.Request, user *users.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

// not using here
func (app *application) contextGetUser(r *http.Request) *users.User {
	user, ok := r.Context().Value(userContextKey).(*users.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
