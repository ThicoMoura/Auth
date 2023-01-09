package controller

import (
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/gin-gonic/gin"
)

type session struct {
	router  *gin.RouterGroup
	service service.Service
}

func (controller session) Setup() {
	controller.router.GET("/")
	controller.router.POST("/")
	controller.router.GET("/:id/")
	controller.router.PUT("/:id/")
	controller.router.DELETE("/:id/")
}

func NewSession(router *gin.RouterGroup, service service.Service) *session {
	return &session{
		router:  router,
		service: service,
	}
}
