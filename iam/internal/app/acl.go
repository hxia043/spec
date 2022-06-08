package app

import (
	"iam/internal/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "missing Authorization in header",
			})
		}

		if !strings.HasPrefix(token, "Bearer ") {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "authorization data is not start with Bearer",
			})
		}

		jwtToken := token[7:]

		claims, err := jwt.ParseToken(jwtToken)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})

			ctx.Abort()
			return
		}

		if claims == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "token invalid",
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
