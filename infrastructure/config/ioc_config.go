package config

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"my-tracking-list-backend/api/v1/auth"
	"my-tracking-list-backend/api/v1/user"
	"my-tracking-list-backend/core/service"
	"my-tracking-list-backend/infrastructure/middleware"
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
	oauthHandler := oauth.NewOauthHandler()
	authMiddleware := middleware.NewAuthenticationMiddleware(userService, oauthHandler)
	userController := user.NewUserController(userService, authMiddleware)

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

func InitIoCManager(engine *gin.Engine) {
	if c == nil {
		containerOnce.Do(func() {
			c = &container{DatabaseConfig()}
		})
	}

	initEssentialsMiddlewares(engine)

	c.InjectUserController(engine)
	c.InjectAuthController(engine)
}

func initEssentialsMiddlewares(engine *gin.Engine) {
	appErroMiddleware := middleware.NewAppErrorMiddleware()
	panicRecoveryMiddleware := middleware.NewPanicRecoveryMiddleware()

	engine.Use(appErroMiddleware.WithAppError)
	engine.Use(gin.CustomRecovery(panicRecoveryMiddleware.WithPanicRecovery))
}