package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/vincentyeungg/golang-menu-creator/db/sqlc"
)

// Server serves HTTP requests for this application
// Server instance has router and db store
type Server struct {
	store  *db.Queries
	router *gin.Engine
}

func SetupServer(store *db.Queries) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// ingredient routes
	ingredientRoutes := router.Group("/api/ingredients")
	ingredientRoutes.GET("/", server.getIngredient)

	// menu item routes
	menuItemRoutes := router.Group("/api/items")
	menuItemRoutes.GET("/", server.getMenuItem)

	// menu routes
	menuRoutes := router.Group("/api/menu")
	menuRoutes.GET("/", server.getMenu)

	// health check routes
	healthRoutes := router.Group("/api/health")
	healthRoutes.GET("/status", server.healthCheck)

	server.router = router

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// start server instance
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
