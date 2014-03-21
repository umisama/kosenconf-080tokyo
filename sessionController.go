package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
)

type SessionController struct{}

// GET /api/session
func (ctrl *SessionController) ReadMany(ctx context.Context) (err error) {
	session, _ := session_store.Get(ctx.HttpRequest(), session_name)
	if session.Values["username"] != nil {
		return goweb.API.RespondWithError(ctx, http.StatusBadRequest, "has session")
	}

	// check form value(not empty?)
	id := ctx.FormValue("id")
	passwd := ctx.FormValue("password")
	if id == "" || passwd == "" {
		return goweb.API.RespondWithError(ctx, http.StatusBadRequest, "not enough query")
	}

	// get userdata from database
	user, err := GetUser(id, passwd)
	if err != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, "error on database")
	}

	if user == nil {
		return goweb.API.RespondWithError(ctx, http.StatusUnauthorized, "user not found")
	}

	session.Values["username"] = user.Name
	session.Save(ctx.HttpRequest(), ctx.HttpResponseWriter())
	return goweb.API.RespondWithData(ctx, nil)
}

// DELETE /api/session
func (ctrl *SessionController) DeleteMany(ctx context.Context) (err error) {
	session, _ := session_store.Get(ctx.HttpRequest(), session_name)

	// clear user's session data
	session.Values["username"] = nil

	session.Save(ctx.HttpRequest(), ctx.HttpResponseWriter())
	return goweb.API.RespondWithData(ctx, nil)
}
