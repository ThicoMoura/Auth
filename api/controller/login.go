package controller

import (
	"net/http"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/ThicoMoura/Auth/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type login struct {
	router  *gin.RouterGroup
	service service.Service
}

func (controller login) Setup() {
	controller.router.POST("/login", controller.login)
}

func (controller login) login(ctx *gin.Context) {
	var req model.Login

	ctx.ShouldBindJSON(&req)

	if validate := util.NewValidate().Validator(req); validate != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Code":    http.StatusBadRequest,
			"Message": validate,
		})
		return
	}

	user, err := controller.service.Get(ctx, req)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "Email incorret")
			return
		}
		if err, ok := err.(*pgconn.PgError); ok {
			switch err.Code {
			}
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	if user.Get("Pass").(string) != req.Pass {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Pass incorret")
		return
	}

	if !user.Get("Active").(bool) {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "Login disable")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func NewLogin(router *gin.RouterGroup, service service.Service) *login {
	return &login{
		router:  router,
		service: service,
	}
}
