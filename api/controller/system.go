package controller

import (
	"net/http"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/gin-gonic/gin"
)

type system struct {
	router  *gin.RouterGroup
	service service.Service
}

func (controller system) Setup() {
	controller.router.GET("/", controller.list)
	controller.router.POST("/", controller.new)
	controller.router.GET("/:id/", controller.get)
	controller.router.PUT("/:id/", controller.update)
	controller.router.DELETE("/:id/", controller.delete)
	controller.router.GET("/search/", controller.find)
}

func (controller system) list(ctx *gin.Context) {
	var req model.List
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if req.PageID != nil && req.PageSize != nil {
		list, err := controller.service.List(ctx, req)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, list)
		return
	}

	list, err := controller.service.List(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, list)
}

func (controller system) new(ctx *gin.Context) {
	var req model.NewSy
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	system, err := controller.service.New(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, system)
}

func (controller system) get(ctx *gin.Context) {
	var req model.Id
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	system, err := controller.service.Get(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, system)
}

func (controller system) update(ctx *gin.Context) {
	var req model.UpdateSy
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	system, err := controller.service.Update(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, system)
}

func (controller system) delete(ctx *gin.Context) {
	var req model.Id
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	system, err := controller.service.Delete(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, system)
}

func (controller system) find(ctx *gin.Context) {
	var req model.FindSy
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	system, err := controller.service.Find(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, system)
}

func NewSystem(router *gin.RouterGroup, service service.Service) *system {
	return &system{
		router:  router,
		service: service,
	}
}
