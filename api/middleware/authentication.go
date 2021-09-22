package middleware

import "github.com/gin-gonic/gin"

type Authentication interface {
	WithAuthentication(ctx *gin.Context)
}
