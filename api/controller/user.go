package controller

import (
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/gin-gonic/gin"
)

type user struct {
	router  *gin.RouterGroup
	service service.Service
}

func (controller user) Setup() {
	controller.router.GET("/")
	controller.router.POST("/")
	controller.router.GET("/:id/")
	controller.router.PUT("/:id/")
	controller.router.DELETE("/:id/")
	controller.router.GET("/profile/")
	controller.router.POST("/access/")
	controller.router.PUT("/access/")
	controller.router.DELETE("/access/")
}

func NewUser(router *gin.RouterGroup, service service.Service) *user {
	return &user{
		router:  router,
		service: service,
	}
}
