package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"vincent.com/todo/controllers"
	"vincent.com/todo/service/auth"
	"vincent.com/todo/service/logger"
	"vincent.com/todo/service/tracing"
)

var log = logger.Logger

func main() {
	// gin.SetMode(gin.ReleaseMode)
	tracing.Init("todo-backend-service")
	defer tracing.Closer.Close()

	r := gin.Default()
	authConn := auth.InitAuthRPC(tracing.Tracer)
	defer authConn.Close()
	r.POST("/api/auth/login", controllers.LoginHandler)
	r.POST("/api/auth/rpc/login", controllers.RPCLoginHandler)
	r.GET("/healthz", controllers.HealthCheckHandler)
	r.GET("/boom", controllers.BoomHandler)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("launch server fail:", zap.String("err", err.Error()))
	}
}
