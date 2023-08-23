package api

import "github.com/gin-gonic/gin"

// Server serves HTTP requests for this application
// Server instance has router and db store
type Server struct {
	router *gin.Engine
}

func SetupServer() *Server {
	server := &Server{}
	router := gin.Default()

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

// start server instance
func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
