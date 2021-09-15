package ioc

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"my-tracking-list-backend/api/v1/auth"
	"my-tracking-list-backend/api/v1/user"
	"my-tracking-list-backend/core/service"
	"my-tracking-list-backend/infrastructure/config"
	"my-tracking-list-backend/infrastructure/oauth"
	db "my-tracking-list-backend/infrastructure/repository"
	"sync"
)

type container struct {
	db *mongo.Database
}

func (c *container) InjectUserController(engine *gin.Engine) {
	userRepository := db.NewUserRepository(c.db)
	userService := service.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	userController.InitRoutes(engine)
}

func (c *container) InjectAuthController(engine *gin.Engine) {
	userRepository := db.NewUserRepository(c.db)
	userService := service.NewUserService(userRepository)

	oauthHandler := oauth.NewOauthHandler()
	authService := service.NewAuthService(oauthHandler, userService)
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
