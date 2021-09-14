package middleware

import (
	"my-tracking-list-backend/core/app_error"
	"fmt"
	"github.com/gin-gonic/gin"
)

func HandleAppError(c *gin.Context) {
	c.Next()
	err := c.Errors.Last()

	if err != nil {
		appError, ok := err.Err.(*app_error.AppError)
		if ok {
			// todo: Usar uma lib de log, tipo uber-go/zap
			fmt.Printf("\nAn error has occured: \t%v\n", appError)
			c.AbortWithStatusJSON(appError.StatusCode, appError)
		} else {
			c.AbortWithStatusJSON(500, app_error.ThrowInternalServerError("An unmapped error has occurred : \t%v\n", err))
		}
	}
}
