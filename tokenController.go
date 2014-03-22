package main

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"github.com/yosida95/random"
)

type TokenController struct{}

func (ctrl *TokenController) ReadMany(ctx context.Context) (err error) {
	t, _ := generateRandomString(tokenlength)
	session, _ := session_store.Get(ctx.HttpRequest(), session_name)
	session.Values["token"] = t

	session.Save(ctx.HttpRequest(), ctx.HttpResponseWriter())
	return goweb.API.RespondWithData(ctx, t)
}

func generateRandomString(length int) (str string, err error) {
	r, err := random.Ascii(random.LOWER | random.UPPER)
	if err != nil {
		return
	}

	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		r, _, _ := r.ReadRune()
		runes[i] = r
	}

	str = string(runes)
	return
}
