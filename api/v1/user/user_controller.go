package user

import (
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/core/ports/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	controller.create()
	controller.findOne()
}

func (controller UserController) create() {
	controller.router.POST("", func(c *gin.Context) {
		var body RequestUser
		if err := c.ShouldBindJSON(&body); err != nil {
			c.Error(fmt.Errorf("erro ao deserializar json do request: %w", err))
			return
		}

		userSaved, err := controller.service.SaveUser(requestUserToUser(body))
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, userSaved)
	})
}

func (controller UserController) findOne() {
	controller.router.GET("/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		userFound, err := controller.service.FindById(id)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, userFound)
	})
}

func requestUserToUser(user RequestUser) domain.User {
	return domain.User{
		ID:       0,
		Email:    user.Email,
		Password: user.Password,
	}
}
