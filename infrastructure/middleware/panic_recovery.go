package middleware

import (
	"my-tracking-list-backend/core/app_error"
	"github.com/gin-gonic/gin"
)

func HandlePanicRecovery(c *gin.Context, _ interface{}) {
	interServerErr := app_error.ThrowInternalServerError("a panic has occurred", nil)
	c.AbortWithStatusJSON(interServerErr.StatusCode, interServerErr)
}
