package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vincentyeungg/golang-menu-creator/db/sqlc"
)

type createIngredientRequest struct {
	Name        string `json:"name" binding:"required"`
	BrandName   string `json:"brand_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
}

func (server *Server) createIngredient(ctx *gin.Context) {
	var req createIngredientRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateIngredientParams{
		Name:        req.Name,
		BrandName:   req.BrandName,
		Description: req.Description,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}

	ingredient, err := server.store.CreateIngredient(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, ingredient)
}

type getIngredientRequest struct {
	ID     int32  `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (server *Server) getIngredient(ctx *gin.Context) {
	var req getIngredientRequest

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

type updateIngredientRequest struct {
	Name        string `json:"name" binding:"required"`
	BrandName   string `json:"brand_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ID          int32  `json:"id" binding:"required"`
}

func (server *Server) updateIngredient(ctx *gin.Context) {
	var req updateIngredientRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateIngredientParams{
		Name:        req.Name,
		BrandName:   req.BrandName,
		Description: req.Description,
		ID:          req.ID,
	}

	ingredient, err := server.store.UpdateIngredient(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ingredient)
}

type deleteIngredientRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) deleteIngredient(ctx *gin.Context) {
	var req deleteIngredientRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteIngredient(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type getAllActiveIngredientsRequest struct {
	PageID   int32 `json:"page_id" binding:"required"`
	PageSize int32 `json:"page_size" binding:"required"`
}

func (server *Server) getAllActiveIngredients(ctx *gin.Context) {
	var req getAllActiveIngredientsRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllActiveIngredientsParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	ingredients, err := server.store.GetAllActiveIngredients(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ingredients)
}

type GetAllIngredientsRequest struct {
	PageID   int32 `json:"page_id" binding:"required"`
	PageSize int32 `json:"page_size" binding:"required"`
}

func (server *Server) getAllIngredients(ctx *gin.Context) {
	var req GetAllIngredientsRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllIngredientParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	ingredients, err := server.store.GetAllIngredient(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ingredients)
}
