package controller

import (
	"net/http"
	"time"

	"github.com/ThicoMoura/Auth/api/model"
	"github.com/ThicoMoura/Auth/api/service"
	"github.com/ThicoMoura/Auth/token"
	"github.com/ThicoMoura/Auth/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type login struct {
	router   *gin.RouterGroup
	service  service.Service
	maker    token.Maker
	duration time.Duration
}

func (controller login) Setup() {
	controller.router.POST("/login/", controller.login)
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

	token, _, err := controller.maker.New(user.Get("Email").(string), controller.duration)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Code":    http.StatusOK,
		"Message": token,
	})
}

func NewLogin(router *gin.RouterGroup, service service.Service, maker token.Maker, duration time.Duration) *login {
	return &login{
		router:   router,
		service:  service,
		maker:    maker,
		duration: duration,
	}
}
