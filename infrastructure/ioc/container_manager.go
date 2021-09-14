package ioc

import (
	"crud-test/api/v1/auth"
	"crud-test/api/v1/user"
	"crud-test/core/service"
	"crud-test/infrastructure/config"
	"crud-test/infrastructure/oauth"
	db "crud-test/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sync"
)

type container struct {
	db *gorm.DB
}

func (c *container) InjectUserController(engine *gin.Engine) {
	userRepository := db.NewUserRepository(c.db)
	userService := service.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	userController.InitRoutes(engine)
}

func (c *container) InjectAuthController(engine *gin.Engine) {
	oauthHandler := oauth.NewOauthHandler()
	authService := service.NewAuthService(oauthHandler)
	authController := auth.NewAuthController(authService)

	authController.InitRoutes(engine)
}

var (
	c             *container
	containerOnce sync.Once
)

func InitContainerManager(engine *gin.Engine) {
	if c == nil {
		containerOnce.Do(func() {
			c = &container{config.DatabaseConfig()}
		})
	}

	c.InjectUserController(engine)
	c.InjectAuthController(engine)
}
