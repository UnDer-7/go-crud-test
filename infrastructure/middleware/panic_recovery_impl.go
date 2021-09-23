package middleware

import (
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/api/middleware"
	"my-tracking-list-backend/core/app_error"
)

type PanicRecoveryImpl struct{}

func NewPanicRecoveryMiddleware() middleware.PanicRecovery {
	return &PanicRecoveryImpl{}
}

func (_ PanicRecoveryImpl) WithPanicRecovery(ctx *gin.Context, _ interface{}) {
	interServerErr := app_error.ThrowInternalServerError("a panic has occurred", nil)
	ctx.AbortWithStatusJSON(interServerErr.StatusCode, interServerErr)
}
