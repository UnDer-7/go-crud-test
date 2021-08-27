package controller

import (
	"crud-test/controller/dto"
	"crud-test/model"
	"crud-test/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	Service service.UserService
	router *gin.RouterGroup
}

func (controller UserController) InitRoutes(engine *gin.Engine)  {
	v1 := engine.Group("v1")
	router := v1.Group("/users")

	controller.router = router

	controller.create()
}

func (controller *UserController) create() {
	controller.router.POST("", func(c *gin.Context) {
		var body dto.UserRequest
		if err := c.ShouldBindJSON(&body); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{"app_error": err.Error()})
			return
		}

		userSaved, err := controller.Service.SaveUser(model.User{
			Email:    body.Email,
			Password: body.Password,
		})

		if err != nil {
			fmt.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"app_error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, userSaved)
	})
}
