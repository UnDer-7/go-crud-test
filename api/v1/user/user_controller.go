package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/api/middleware"
	"my-tracking-list-backend/core/help"
	"my-tracking-list-backend/core/ports/driver"
	"net/http"
)

type UserController struct {
	authMiddleware middleware.Authentication
	service        driver.UserService
	router         *gin.RouterGroup
}

func NewUserController(service driver.UserService, authMiddleware middleware.Authentication) *UserController {
	return &UserController{
		service:        service,
		authMiddleware: authMiddleware,
	}
}

func (controller UserController) InitRoutes(engine *gin.Engine) {
	v1 := engine.Group("v1")
	router := v1.Group("/users")
	controller.router = router

	controller.router.Use(controller.authMiddleware.WithAuthentication)

	controller.findOne()
}

func (controller UserController) findOne() {
	controller.router.GET("/email/:email", func(c *gin.Context) {
		email := c.Param("email")

		// todo: Criar get/set para pegar valores do context
		// https://medium.com/@matryer/context-keys-in-go-5312346a868d
		tmp, err := help.GetCurrentUserEmail(c)
		if err != nil {
			c.Error(err)
			return
		}

		fmt.Printf("\n$$$$$$$$EMAIL FROM CTX: %s\n$$$$$", tmp)
		userFound, err := controller.service.FindByEmail(email)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, userToResponseUser(userFound))
	})
}
