package authtransport

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (t *httpTransport) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "done",
	})
}
