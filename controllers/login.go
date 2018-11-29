package controllers

import (
	"go.uber.org/zap"
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
		log.Sugar().Error("LoginHandler", "receive params err", login)
		res.Err(c, err.Error())
		return
	}
	user := models.User{
		Name: login.User,
	}
	// 此处模拟检查用户,获取uid过程
	user.NewUID()
	log.Info("new uid", zap.String("uid", user.ID))
	err := user.NewToken()
	if err != nil {
		log.Error("new token", zap.String("err", err.Error()))
		res.Err(c, err.Error())
		return
	}
	log.Sugar().Infow("LoginHandler res", user)
	res.JSON(c, &user)
}
