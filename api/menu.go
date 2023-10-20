package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/vincentyeungg/golang-menu-creator/db/sqlc"
)

type createMenuRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
	UpdatedBy   string `json:"updated_by" binding:"required"`
}

func (server *Server) createMenu(ctx *gin.Context) {
	var req createMenuRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateMenuParams{
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.UpdatedBy,
	}

	menu, err := server.store.CreateMenu(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, menu)
}

type getMenuRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) getMenu(ctx *gin.Context) {
	var req getMenuRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	menu, err := server.store.GetMenu(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

type updateMenuRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ID          int32  `json:"id" binding:"required"`
}

func (server *Server) updateMenu(ctx *gin.Context) {
	var req updateMenuRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateMenuParams{
		Name:        req.Name,
		Description: req.Description,
		ID:          req.ID,
	}

	menu, err := server.store.UpdateMenu(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menu)
}

type deleteMenuRequest struct {
	ID int32 `json:"id" binding:"required"`
}

func (server *Server) deleteMenu(ctx *gin.Context) {
	var req deleteMenuRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteMenu(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

type getAllMenusRequest struct {
	PageID   int32 `json:"page_id" binding:"required"`
	PageSize int32 `json:"page_size" binding:"required"`
}

func (server *Server) getAllMenus(ctx *gin.Context) {
	var req getAllMenusRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllMenusParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	menus, err := server.store.GetAllMenus(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menus)
}

type getAllActiveMenusRequest struct {
	PageID   int32 `json:"page_id" binding:"required"`
	PageSize int32 `json:"page_size" binding:"required"`
}

func (server *Server) getAllActiveMenus(ctx *gin.Context) {
	var req getAllActiveMenusRequest

	// verify input req arguments
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetAllActiveMenusParams{
		Limit:  req.PageID,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	menus, err := server.store.GetAllActiveMenus(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, menus)
}
