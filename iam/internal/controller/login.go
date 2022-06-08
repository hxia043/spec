package controller

import (
	"iam/internal/models"
	"iam/internal/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

func Login(ctx *gin.Context) {
	loginUser := new(LoginUser)
	if err := ctx.ShouldBind(loginUser); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	if err := models.Login(loginUser.Name, loginUser.Password); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	token, err := jwt.GenerateToken(loginUser.Name)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "login success",
		"token": token,
	})
}
