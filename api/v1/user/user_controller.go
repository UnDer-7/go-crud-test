package user

import (
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/core/ports/driver"
	"net/http"
)

type UserController struct {
	service driver.UserService
	router  *gin.RouterGroup
}

func NewUserController(service driver.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (controller UserController) InitRoutes(engine *gin.Engine) {
	v1 := engine.Group("v1")
	router := v1.Group("/users")

	controller.router = router

	controller.findOne()
}

func (controller UserController) findOne() {
	controller.router.GET("/email/:email", func(c *gin.Context) {
		email := c.Param("email")

		userFound, err := controller.service.FindByEmail(email)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, userToResponseUser(userFound))
	})
}

