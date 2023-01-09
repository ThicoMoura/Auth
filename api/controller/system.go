package controller

import (
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/gin-gonic/gin"
)

type system struct {
	router  *gin.RouterGroup
	service service.Service
}

func (controller system) Setup() {
	controller.router.GET("/")
	controller.router.POST("/")
	controller.router.GET("/:id/")
	controller.router.PUT("/:id/")
	controller.router.DELETE("/:id/")
}

func NewSystem(router *gin.RouterGroup, service service.Service) *system {
	return &system{
		router:  router,
		service: service,
	}
}
