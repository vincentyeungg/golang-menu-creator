package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getMenu(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "get menu api",
	})

	return
}