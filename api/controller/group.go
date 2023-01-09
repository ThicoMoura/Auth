package controller

import (
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/gin-gonic/gin"
)

type group struct {
	router  *gin.RouterGroup
	service service.Service
}

func (controller group) Setup() {
	controller.router.GET("/")
	controller.router.POST("/")
	controller.router.GET("/:id/")
	controller.router.PUT("/:id/")
	controller.router.DELETE("/:id/")
	controller.router.POST("/access/")
	controller.router.PUT("/access/")
	controller.router.DELETE("/access/")
}

func NewGroup(router *gin.RouterGroup, service service.Service) *group {
	return &group{
		router:  router,
		service: service,
	}
}
