package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kishpower/simplebank/db/sqlc"
)

// server serves http requests for our banking application
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// creates a new Http server and setup routes
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//adds routes to server
	router.POST("/accounts", server.createAccount)
	// router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	// todo : update and delete
	server.router = router
	return server
}

// starts the server on the given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}

}
