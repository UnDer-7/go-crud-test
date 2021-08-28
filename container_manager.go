package main

import (
	"crud-test/api/v1/user"
	"crud-test/core/domain"
	"crud-test/core/service"
	db "crud-test/infrastructure/database"
	"github.com/gin-gonic/gin"
	"sync"
)

type container struct{}

func (c *container) InjectUserController(engine *gin.Engine) {
	database := make(map[int]domain.User)

	userRepository := &db.UserRepositoryImpl{Database: database}
	userService := &service.UserServiceImpl{Repository: userRepository}
	userController := &user.UserController{Service: userService}

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
