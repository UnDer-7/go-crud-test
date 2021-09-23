package auth

import (
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/core/app_error"
	"my-tracking-list-backend/core/ports/driver"
)

type Controller struct {
	service driver.AuthService
	router  *gin.RouterGroup
}

func NewAuthController(service driver.AuthService) *Controller {
	return &Controller{service: service}
}

func (c Controller) InitRoutes(engine *gin.Engine) {
	v1 := engine.Group("v1")
	router := v1.Group("/auth")

	c.router = router

	c.signIn()
}

func (c Controller) signIn() {
	c.router.POST("/sign-in/google", func(ctx *gin.Context) {
		var body RequestToken

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.Error(app_error.ThrowInternalServerError("Erro ao deserializar token", err))
			return
		}

		usr, err := c.service.SignIn(body.Token)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.JSON(200, userToResponseUser(usr))
	})
}
