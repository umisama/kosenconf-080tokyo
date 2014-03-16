package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
)

type UserController struct{}

// POST /api/user
func (ctrl *UserController) Create(ctx context.Context) (err error) {
	id := ctx.FormValue("id")
	screen_name := ctx.FormValue("screen_name")
	passwd := ctx.FormValue("password")
	if id == "" || screen_name == "" || passwd == "" {
		return goweb.API.RespondWithError(ctx, http.StatusBadRequest, "not enough query.")
	}

	err = CreateUser(id, screen_name, passwd)
	if err != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, "error on database")
	}

	return goweb.API.RespondWithData(ctx, nil)
}
