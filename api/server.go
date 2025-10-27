package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/junasliang/go_simplebank/db/sqlc"
)

// Server serve HTTP requests for banking
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Create HTTP server and setup router
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// define apis
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
