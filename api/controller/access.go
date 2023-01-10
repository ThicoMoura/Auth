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
	controller.router.GET("/", controller.list)
	controller.router.POST("/", controller.new)
	controller.router.GET("/:id/", controller.get)
	controller.router.PUT("/:id/", controller.update)
	controller.router.DELETE("/:id/", controller.delete)
	controller.router.POST("/:id/group/", controller.insertGroups)
	controller.router.PUT("/:id/group/", controller.replaceGroups)
	controller.router.DELETE("/:id/group/", controller.deleteGroups)
	controller.router.POST("/:id/user/", controller.insertUsers)
	controller.router.PUT("/:id/user/", controller.replaceUsers)
	controller.router.DELETE("/:id/user/", controller.deleteUsers)
}

func (controller access) list(ctx *gin.Context) {}

func (controller access) new(ctx *gin.Context) {}

func (controller access) get(ctx *gin.Context) {
	var req model.Id
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	access, err := controller.service.Get(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, access)
}

func (controller access) update(ctx *gin.Context) {}

func (controller access) delete(ctx *gin.Context) {}

func (controller access) insertGroups(ctx *gin.Context) {}

func (controller access) replaceGroups(ctx *gin.Context) {}

func (controller access) deleteGroups(ctx *gin.Context) {}

func (controller access) insertUsers(ctx *gin.Context) {}

func (controller access) replaceUsers(ctx *gin.Context) {}

func (controller access) deleteUsers(ctx *gin.Context) {}

func NewAccess(router *gin.RouterGroup, service service.Service) *access {
	return &access{
		router:  router,
		service: service,
	}
}
