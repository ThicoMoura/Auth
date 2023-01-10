package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ThicoMoura/Auth/token"
	"github.com/gin-gonic/gin"
)

type md struct {
	maker token.Maker
}

func (middle md) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Code":    http.StatusUnauthorized,
				"Message": "Authorization header is not provided",
			})
			return
		}

		fields := strings.Fields(header)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Code":    http.StatusUnauthorized,
				"Message": "Invalid authorization header format",
			})
			return
		}

		if authorizationType := strings.ToLower(fields[0]); authorizationType != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Code":    http.StatusUnauthorized,
				"Message": fmt.Sprintf("unsupported authorization type %s", authorizationType),
			})
			return
		}

		payload, err := middle.maker.Valid(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Code":    http.StatusUnauthorized,
				"Message": err,
			})
			return
		}

		ctx.Set("payload", payload)
		ctx.Next()
	}
}

func NewMiddleware(maker token.Maker) *md {
	return &md{
		maker: maker,
	}
}
