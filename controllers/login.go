package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"vincent.com/todo/models"
	"vincent.com/todo/service/res"
	"vincent.com/todo/service/tracing"
)

// LoginHandler - handler login api
func LoginHandler(c *gin.Context) {
	tracer := tracing.Tracer
	span := tracer.StartSpan("LoginHandler")
	defer span.Finish()
	type LoginReq struct {
		User     string `json:"username"  binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var login LoginReq
	if err := c.ShouldBindJSON(&login); err != nil {
		span.LogKV("event", "bind", "err", err)
		log.Sugar().Error("LoginHandler", "receive params err", login)
		res.Err(c, err.Error())
		return
	}
	user := models.User{
		Name: login.User,
	}
	ctx := opentracing.ContextWithSpan(context.Background(), span)
	// 此处模拟检查用户,获取uid过程
	user.NewUID(ctx)
	log.Info("new uid", zap.String("uid", user.ID))
	err := user.NewToken(ctx)
	if err != nil {
		log.Error("new token", zap.String("err", err.Error()))
		res.Err(c, err.Error())
		return
	}
	log.Sugar().Infow("LoginHandler res", "userInfo", user)
	res.JSON(c, &user)
}
