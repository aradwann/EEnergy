package api

import (
	"fmt"

	db "github.com/aradwann/eenergy/db/store"
	"github.com/aradwann/eenergy/token"
	"github.com/aradwann/eenergy/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HTTP request for out banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {

	tokenMaker, err := token.NewPASETOMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}
	server := &Server{store: store, tokenMaker: tokenMaker, config: config}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("unit", validUnit)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
