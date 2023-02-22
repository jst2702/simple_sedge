package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"simplesedge.com/feed/pkg/db"
)

// serves all HTTP requests for our feed service
type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("password", validPassword)
	}

	router.POST("/users", server.createUser)
	router.GET("/users/:email", server.getUser)
	router.GET("/documents/:guid", server.getDocument)
	router.GET("/documents", server.listDocuments)

	// add routes to router
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
