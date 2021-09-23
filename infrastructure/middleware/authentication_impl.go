package middleware

import (
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/api/middleware"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/ports/driven"
	"my-tracking-list-backend/core/ports/driver"
)

type AuthenticationImpl struct {
	userService  driver.UserService
	oauthHandler driven.OauthHandler
}

func NewAuthenticationMiddleware(userService driver.UserService, oauthHandler driven.OauthHandler) middleware.Authentication {
	return &AuthenticationImpl{
		userService:  userService,
		oauthHandler: oauthHandler,
	}
}

func (a AuthenticationImpl) WithAuthentication(ctx *gin.Context) {
	whenAnErrorOccurs := func(err error) {
		ctx.Error(err)
		ctx.Abort()
	}

	token := ctx.GetHeader("Authorization")
	// todo: vai da merda quando fizer com facebook login
	gToken, err := a.oauthHandler.DecodeGoogleToken(token)
	if err != nil {
		whenAnErrorOccurs(err)
		return
	}

	exists, err := a.userService.UserExistes(gToken.Email)
	if err != nil {
		whenAnErrorOccurs(err)
		return
	}

	if !exists {
		whenAnErrorOccurs(app_error.ThrowNotFoundError(
			"Usuario nao encontrado",
			"Usuario com email informado nao foi cadastrado",
			nil),
		)
		return
	}

	// todo: Criar um help pra setar e pegar os values do context
	// https://medium.com/@matryer/context-keys-in-go-5312346a868d
	ctx.Set("UserEmail", gToken.Email)
	ctx.Next()
}
