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
	controller.router.GET("/", controller.list)
	controller.router.POST("/", controller.new)
	controller.router.GET("/:id/", controller.get)
	controller.router.PUT("/:id/", controller.update)
	controller.router.DELETE("/:id/", controller.delete)
	controller.router.GET("/profile/", controller.profile)
	controller.router.POST("/access/", controller.insertAccess)
	controller.router.PUT("/access/", controller.replaceAccess)
	controller.router.DELETE("/access/", controller.deleteAccess)
}

func (controller user) list(ctx *gin.Context)          {}
func (controller user) new(ctx *gin.Context)           {}
func (controller user) get(ctx *gin.Context)           {}
func (controller user) update(ctx *gin.Context)        {}
func (controller user) delete(ctx *gin.Context)        {}
func (controller user) profile(ctx *gin.Context)       {}
func (controller user) insertAccess(ctx *gin.Context)  {}
func (controller user) replaceAccess(ctx *gin.Context) {}
func (controller user) deleteAccess(ctx *gin.Context)  {}

func NewUser(router *gin.RouterGroup, service service.Service) *user {
	return &user{
		router:  router,
		service: service,
	}
}
