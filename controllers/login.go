package controllers

import (
	"vincent.com/todo/service/res"

	"github.com/gin-gonic/gin"
	"vincent.com/todo/models"
)

// LoginHandler - handler login api
func LoginHandler(c *gin.Context) {
	type LoginReq struct {
		User     string `json:"username"  binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var login LoginReq
	if err := c.ShouldBindJSON(&login); err != nil {
		res.Err(c, err.Error())
		return
	}
	user := models.User{
		Name: login.User,
	}
	// 此处模拟检查用户,获取uid过程
	user.NewUID()
	err := user.NewToken()
	if err != nil {
		res.Err(c, err.Error())
		return
	}

	res.JSON(c, &user)
}
