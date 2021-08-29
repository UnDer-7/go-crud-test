package middleware

import (
	"crud-test/core/app_error"
	"github.com/gin-gonic/gin"
)

func HandleAppError(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()

	if err != nil {
		appError, ok := err.Err.(*app_error.AppError)
		if ok {
			c.AbortWithStatusJSON(appError.StatusCode, appError)
		} else {
			c.AbortWithStatusJSON(500, app_error.ThrowInternalServerError("An unmapped error has occurred : \t%v\n", err))
		}
	}
}
