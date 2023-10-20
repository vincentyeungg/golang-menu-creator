package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vincentyeungg/golang-menu-creator/db/sqlc"
)

type createMenuItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Status      string `json:"status" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
}

func (server *Server) createMenuItem(ctx *gin.Context) {
	var req createMenuItemRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMenuItemParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}

	menuItem, err := server.store.CreateMenuItem(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, menuItem)
}

type getMenuItemRequest struct {
	ID     int32  `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (server *Server) getMenuItem(ctx *gin.Context) {
	var req getMenuItemRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetMenuItemParams{
		ID:     req.ID,
		Status: req.Status,
	}

	// query db for menuItem with id and status
	menuItem, err := server.store.GetMenuItem(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menuItem)
}

type updateMenuItemRequest struct {
	Name        string `json:"name" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	ID          int32  `json:"id" binding:"required"`
}

func (server *Server) updateMenuItem(ctx *gin.Context) {
	var req updateMenuItemRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMenuItemParams{
		Name:        req.Name,
		Price:       req.Price,
		Description: req.Description,
		ID:          req.ID,
	}

	menuItem, err := server.store.UpdateMenuItem(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menuItem)
}

type deleteMenuItemRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) deleteMenuItem(ctx *gin.Context) {
	var req deleteMenuItemRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteMenuItem(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type getAllMenuItemsRequest struct {
	PageID   int32 `json:"page_id" binding:"required"`
	PageSize int32 `json:"page_size" binding:"required"`
}

func (server *Server) getAllMenuItems(ctx *gin.Context) {
	var req getAllMenuItemsRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllMenuItemsParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	menuItems, err := server.store.GetAllMenuItems(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menuItems)
}

type getAllActiveMenuItemsRequest struct {
	PageID   int32 `json:"page_id" binding:"required"`
	PageSize int32 `json:"page_size" binding:"required"`
}

func (server *Server) getAllActiveMenuItems(ctx *gin.Context) {
	var req getAllActiveMenuItemsRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllActiveItemsParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	menuItems, err := server.store.GetAllActiveItems(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menuItems)
}
