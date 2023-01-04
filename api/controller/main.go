package controller

import (
	"github.com/ThicoMoura/Auth/db/repository"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *repository.Repository
	router *gin.Engine
}

func (server Server) setup() {
	server.router.SetTrustedProxies([]string{})
}

func (server Server) Start(addr string) {
	server.setup()
	server.router.Run(addr)
}

func NewServer(store *repository.Repository) *Server {
	return &Server{
		store:  store,
		router: gin.Default(),
	}
}
