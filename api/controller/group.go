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
	controller.router.GET("/", controller.list)
	controller.router.POST("/", controller.new)
	controller.router.GET("/:id/", controller.get)
	controller.router.PUT("/:id/", controller.update)
	controller.router.DELETE("/:id/", controller.delete)
	controller.router.POST("/access/", controller.insertAccess)
	controller.router.PUT("/access/", controller.replaceAccess)
	controller.router.DELETE("/access/", controller.deleteAccess)
}

func (controller group) list(ctx *gin.Context)          {}
func (controller group) new(ctx *gin.Context)           {}
func (controller group) get(ctx *gin.Context)           {}
func (controller group) update(ctx *gin.Context)        {}
func (controller group) delete(ctx *gin.Context)        {}
func (controller group) insertAccess(ctx *gin.Context)  {}
func (controller group) replaceAccess(ctx *gin.Context) {}
func (controller group) deleteAccess(ctx *gin.Context)  {}

func NewGroup(router *gin.RouterGroup, service service.Service) *group {
	return &group{
		router:  router,
		service: service,
	}
}
