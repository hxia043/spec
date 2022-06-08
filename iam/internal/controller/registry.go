package controller

import (
	"fmt"
	"iam/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

func Registry(ctx *gin.Context) {
	// config username and password according to configmap of default_user
	user := new(User)

	if err := ctx.ShouldBind(user); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})

		return
	}

	fmt.Println(user.Name, user.Password)

	// save username and password
	if err := models.InsertUser(user.Name, user.Password); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "registry successful",
	})
}
