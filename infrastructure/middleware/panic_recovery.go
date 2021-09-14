package middleware

import (
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/core/app_error"
)

func HandlePanicRecovery(c *gin.Context, _ interface{}) {
	interServerErr := app_error.ThrowInternalServerError("a panic has occurred", nil)
	c.AbortWithStatusJSON(interServerErr.StatusCode, interServerErr)
}
