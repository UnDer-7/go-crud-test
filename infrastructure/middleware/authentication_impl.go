package middleware

import (
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/api/middleware"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/help"
	"my-tracking-list-backend/core/ports/driven"
	"my-tracking-list-backend/core/ports/driver"
)

type authentication struct {
	userService  driver.UserService
	oauthHandler driven.OauthHandler
}

func NewAuthenticationMiddleware(userService driver.UserService, oauthHandler driven.OauthHandler) middleware.Authentication {
	return &authentication{
		userService:  userService,
		oauthHandler: oauthHandler,
	}
}

func (a authentication) WithAuthentication(ctx *gin.Context) {
	whenAnErrorOccurs := func(err error) {
		ctx.Error(err)
		ctx.Abort()
	}

	token := ctx.GetHeader("Authorization")
	// todo: vai da merda quando fizer com facebook login
	gToken, err := a.oauthHandler.DecodeGoogleToken(ctx, token)
	if err != nil {
		whenAnErrorOccurs(err)
		return
	}

	exists, err := a.userService.UserExists(ctx, gToken.Email)
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

	if err := help.SetCurrentUserEmail(ctx, gToken.Email); err != nil {
		whenAnErrorOccurs(err)
		return
	}

	ctx.Next()
}
