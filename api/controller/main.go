package controller

import (
	"net/http"
	"time"

	"github.com/ThicoMoura/Auth/api/service"
	"github.com/ThicoMoura/Auth/db/repository"
	"github.com/ThicoMoura/Auth/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *repository.Repository
	token  token.Maker
	router *gin.Engine
}

func (server Server) setup() {
	server.router.SetTrustedProxies([]string{})

	server.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := server.router.Group("/api/v1")
	md := NewMiddleware(server.token)

	NewLogin(api, service.NewLogin(map[string]repository.IRepository{
		"user":    server.store.Table("user"),
		"session": server.store.Table("session"),
	}), server.token, time.Hour, md.Authentication()).Setup()

	auth := api.Group("/")

	auth.Use(md.Authentication())

	NewAccess(auth.Group("/access"), service.NewAccess(server.store.Table("access"))).Setup()
	NewGroup(auth.Group("/group"), service.NewGroup(server.store.Table("group"))).Setup()
	NewSession(auth.Group("/session"), service.NewSession(server.store.Table("session"))).Setup()
	NewSystem(auth.Group("/system"), service.NewSystem(server.store.Table("system"))).Setup()
	NewUser(auth.Group("/user"), service.NewUser(server.store.Table("user"))).Setup()
}

func (server Server) Start(addr string) *http.Server {
	server.setup()
	return &http.Server{
		Addr:    addr,
		Handler: server.router,
	}
}

func NewServer(store *repository.Repository, token token.Maker) *Server {
	return &Server{
		store:  store,
		token:  token,
		router: gin.Default(),
	}
}
