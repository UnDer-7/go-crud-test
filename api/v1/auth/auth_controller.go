package auth

import (
	"crud-test/core/app_error"
	"crud-test/core/ports/driver"
	"github.com/gin-gonic/gin"
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

	c.login()
}

func (c Controller) login() {
	router := c.router
	router.POST("/login/google", func(ctx *gin.Context) {
		var body RequestToken

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.Error(app_error.ThrowInternalServerError("Erro deserializar token", err))
			return
		}

		if err := c.service.Login(body.Token); err != nil {
			ctx.Error(err)
			return
		}
	})
}
