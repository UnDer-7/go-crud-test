package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSConfig() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://127.0.0.1:3000",
	}

	return cors.New(config)
}
