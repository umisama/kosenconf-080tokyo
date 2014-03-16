package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
	"strconv"
)

type StatusesController struct{}

// POST /api/statuses
func (ctrl *StatusesController) Create(ctx context.Context) (err error) {
	user := getUserNameFromCtx(ctx)
	if user == "" {
		return goweb.API.RespondWithError(ctx, http.StatusUnauthorized, "not enough query")
	}

	text := ctx.FormValue("shout")
	if text == "" {
		return goweb.API.RespondWithError(ctx, http.StatusBadRequest, "not enough query")
	}

	err = CreateStatus(user, text)
	if err != nil {
		logger.Debug(err)
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, "error on db")
	}

	return goweb.API.RespondWithData(ctx, nil)
}

// GET /api/statuses
func (ctrl *StatusesController) ReadMany(ctx context.Context) (err error) {
	count, err := strconv.Atoi(ctx.FormValue("count"))
	if err != nil {
		count = 15
	}

	dat, err := GetStatuses(count)
	if err != nil {
		logger.Debug(err)
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, "error on db")
	}

	return goweb.API.RespondWithData(ctx, dat)
}

func getUserNameFromCtx(ctx context.Context) (name string) {
	session, _ := session_store.Get(ctx.HttpRequest(), session_name)
	user_raw := session.Values["username"]
	if user_raw == nil {
		return ""
	}
	name = user_raw.(string)
	return
}
