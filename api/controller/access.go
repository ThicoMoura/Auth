package controller

import (
	"net/http"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/gin-gonic/gin"
)

type access struct {
	router  *gin.RouterGroup
	service service.Service
}

func (controller access) Setup() {
	controller.router.GET("/")
	controller.router.POST("/")
	controller.router.GET("/:id/", controller.get)
	controller.router.PUT("/:id/")
	controller.router.DELETE("/:id/")
	controller.router.POST("/group/")
	controller.router.PUT("/group/")
	controller.router.DELETE("/group/")
	controller.router.POST("/user/")
	controller.router.PUT("/user/")
	controller.router.DELETE("/user/")
}

func (controller access) get(ctx *gin.Context) {
	var req model.Id
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	access, err := controller.service.Get(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, access)
}

func NewAccess(router *gin.RouterGroup, service service.Service) *access {
	return &access{
		router:  router,
		service: service,
	}
}
