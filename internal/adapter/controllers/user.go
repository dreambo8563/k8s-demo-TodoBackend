package controllers

import (
	"github.com/gin-gonic/gin"
	"vincent.com/todo/internal/adapter/service"
	"vincent.com/todo/internal/pkg/res"
)

//UserInfo - get UserInfo
func UserInfo(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		res.Err(c, "not found token")
		return
	}
	userService := service.InitializeUserCase()
	user, err := userService.GetInfo(c, token)
	if err != nil {
		res.Err(c, err.Error())
		return
	}
	res.JSON(c, user)
}
