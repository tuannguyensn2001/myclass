package authtransport

import (
	"github.com/gin-gonic/gin"
	"myclass/src/app"
	"myclass/src/services/auth/authstruct"
	"net/http"
)

func (t *httpTransport) Register(ctx *gin.Context) {
	var input authstruct.RegisterInput

	if err := ctx.ShouldBind(&input); err != nil {
		panic(app.BadRequestHttpError("data not valid", err))
	}

	err := t.service.Register(ctx.Request.Context(), input)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})
}
