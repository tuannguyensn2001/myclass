package middlewares

import (
	"github.com/gin-gonic/gin"
	"myclass/src/app"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if httpError, ok := err.(*app.HttpError); ok {
				ctx.AbortWithStatusJSON(httpError.StatusCode, httpError)
				return
			}
			httpError := app.InternalHttpError("internal server", err.(error))
			ctx.AbortWithStatusJSON(httpError.StatusCode, httpError)
			return
		}
	}()
	ctx.Next()
}
