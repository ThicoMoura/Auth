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
	controller.router.GET("/", controller.list)
	controller.router.GET("/:id/", controller.get)
	controller.router.DELETE("/:id/", controller.delete)
	controller.router.GET("/search/", controller.find)
}

func (controller session) list(ctx *gin.Context)   {}
func (controller session) get(ctx *gin.Context)    {}
func (controller session) delete(ctx *gin.Context) {}
func (controller session) find(ctx *gin.Context)   {}

func NewSession(router *gin.RouterGroup, service service.Service) *session {
	return &session{
		router:  router,
		service: service,
	}
}
