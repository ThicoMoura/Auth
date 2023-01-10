package controller

import (
	"net/http"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/gin-gonic/gin"
)

type group struct {
	router   *gin.RouterGroup
	gService service.Service
}

func (controller group) Setup() {
	controller.router.GET("/", controller.list)
	controller.router.POST("/", controller.new)
	controller.router.GET("/:id/", controller.get)
	controller.router.PUT("/:id/", controller.update)
	controller.router.DELETE("/:id/", controller.delete)
	controller.router.GET("/search/", controller.find)
	controller.router.POST("/access/", controller.insertAccess)
	controller.router.PUT("/access/", controller.replaceAccess)
	controller.router.DELETE("/access/", controller.deleteAccess)
}

func (controller group) list(ctx *gin.Context) {
	var req model.List
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	group, err := controller.gService.List(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (controller group) new(ctx *gin.Context) {
	var req model.NewG
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	group, err := controller.gService.New(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, group)
}

func (controller group) get(ctx *gin.Context) {
	var req model.Id
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	group, err := controller.gService.Get(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (controller group) update(ctx *gin.Context) {
	var req model.UpdateG
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	group, err := controller.gService.Update(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (controller group) delete(ctx *gin.Context) {
	var req model.Id
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	group, err := controller.gService.Delete(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, group)
}

func (controller group) find(ctx *gin.Context) {}

func (controller group) insertAccess(ctx *gin.Context) {}

func (controller group) replaceAccess(ctx *gin.Context) {}

func (controller group) deleteAccess(ctx *gin.Context) {}

func NewGroup(router *gin.RouterGroup, gService service.Service) *group {
	return &group{
		router:   router,
		gService: gService,
	}
}
