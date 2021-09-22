package middleware

import "github.com/gin-gonic/gin"

type AppError interface {
	WithAppError(ctx *gin.Context)
}
