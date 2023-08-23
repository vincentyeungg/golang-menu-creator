package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "App health is ok!",
	})

	return
}