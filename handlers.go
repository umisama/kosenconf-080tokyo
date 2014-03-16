package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
)

func beforeHandler(c context.Context) (err error) {
	session, _ := session_store.Get(c.HttpRequest(), session_name)
	session.Save(c.HttpRequest(), c.HttpResponseWriter())
	return nil
}

func afterHandler(c context.Context) (err error) {
	logger.Infof("[access] /%s", c.Path().RawPath)
	return nil
}

func apiGetSessionHandler(c context.Context) (err error) {
	return goweb.API.RespondWithError(c, 501, "not implemented")
}
