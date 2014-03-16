package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
	"strconv"
)

type SearchController struct{}

// GET /api/session
func (ctrl *SearchController) ReadMany(ctx context.Context)(err error) {
	query := ctx.FormValue("q")
	if query == "" {
		return goweb.API.RespondWithError(ctx, http.StatusUnauthorized, "not enough query")
	}

	count, err := strconv.Atoi(ctx.FormValue("count"))
	if err != nil {
		count = 10
	}

	statuses, err := SearchStatuses(query, count)
	if err != nil {
		logger.Debug(err)
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, "error on db")
	}

	return goweb.API.RespondWithData(ctx, statuses)
}
