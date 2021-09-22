package middleware

import "github.com/gin-gonic/gin"

type PanicRecovery interface {
	WithPanicRecovery(ctx *gin.Context, _ interface{})
}
