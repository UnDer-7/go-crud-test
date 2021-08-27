package main

import (
	"crud-test/controller"
	"crud-test/repositories"
	"crud-test/service"
	"github.com/gin-gonic/gin"
	"sync"
)

type container struct{}

func (c *container) InjectUserController(engine *gin.Engine) {
	// todo: Config database

	userRepository := &repositories.UserRepository{}
	userService := service.UserService{Repository: userRepository}
	userController := controller.UserController{Service: userService}

	userController.InitRoutes(engine)
}

var (
	c             *container
	containerOnce sync.Once
)

func InitContainerManager(engine *gin.Engine) {
	if c == nil {
		containerOnce.Do(func() {
			c = &container{}
		})
	}

	c.InjectUserController(engine)
}
