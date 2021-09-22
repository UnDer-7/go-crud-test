package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my-tracking-list-backend/api/middleware"
	"my-tracking-list-backend/core/app_error"
)

type AppErrorImpl struct {}

func NewAppErrorMiddleware() middleware.AppError {
	return &AppErrorImpl{}
}

func (a AppErrorImpl) WithAppError(ctx *gin.Context) {
	ctx.Next()
	err := ctx.Errors.Last()

	if err != nil {
		appError, ok := err.Err.(*app_error.AppError)
		if ok {
			// todo: Usar uma lib de log, tipo uber-go/zap
			fmt.Printf("\nAn error has occured: \t%v\n", appError)
			ctx.AbortWithStatusJSON(appError.StatusCode, appError)
		} else {
			ctx.AbortWithStatusJSON(500, app_error.ThrowInternalServerError("An unmapped error has occurred : \t%v\n", err))
		}
	}
}
