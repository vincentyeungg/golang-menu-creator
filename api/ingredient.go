package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vincentyeungg/golang-menu-creator/db/sqlc"
)

type getIngredientRequest struct {
	ID     int32  `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (server *Server) getIngredient(ctx *gin.Context) {
	var req getIngredientRequest
	var err error

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetIngredientParams{
		ID:     req.ID,
		Status: req.Status,
	}

	// query db for ingredient with id and status
	ingredient, err := server.store.GetIngredient(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ingredient)
}
