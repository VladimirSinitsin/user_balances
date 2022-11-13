package api

import (
	db "github.com/VladimirSinitsin/user_balances/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.GET("/account/balance/:id", server.getAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorRequest ...
func errorRequest(err error) gin.H {
	return gin.H{"error": err.Error()}
}
