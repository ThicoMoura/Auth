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

	NewLogin(api, service.NewLogin(map[string]repository.IRepository{
		"user":    server.store.Table("user"),
		"session": server.store.Table("session"),
	}), server.token, time.Hour).Setup()

	auth := api.Group("/")
	md := NewMiddleware(server.token, nil)

	auth.Use(md.Authentication())

	auth.GET("/profile", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Code": http.StatusOK,
			"Message": map[string]string{
				"Name":  "PIPI",
				"Email": "POPO",
			},
		})
	})
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
