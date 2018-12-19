package controllers

import (
	"context"

	"vincent.com/todo/internal/adapter/service"

	"vincent.com/todo/internal/domain/usecase"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"vincent.com/todo/internal/pkg/res"
	"vincent.com/todo/internal/pkg/tracing"
)

// RegisterHandler - handler login api
func RegisterHandler(c *gin.Context) {
	tracer := tracing.Tracer
	span := tracer.StartSpan("RegisterHandler")
	defer span.Finish()
	type LoginReq struct {
		User     string `json:"username"  binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var login LoginReq
	if err := c.ShouldBindJSON(&login); err != nil {
		span.LogKV("event", "bind", "err", err)
		log.Sugar().Error("RegisterHandler", "receive params err", login)
		res.Err(c, err.Error())
		return
	}

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	// 此处模拟检查用户,获取uid过程
	repo := service.NewUserRepository()
	uc := usecase.NewUserUsecase(repo)

	user, token, err := uc.RegisterUser(ctx, login.User, login.Password)

	if err != nil {
		log.Error("RegisterUser", zap.String("err", err.Error()))
		res.Err(c, err.Error())
		return
	}
	log.Sugar().Infow("RegisterHandler res", "userInfo", user)
	type registerRes struct {
		User  *usecase.User `json:"user"`
		Token string        `json:"token"`
	}
	res.JSON(c, &registerRes{
		User:  user,
		Token: token,
	})
}
