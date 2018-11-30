package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"vincent.com/todo/controllers"
	"vincent.com/todo/service/logger"
)

var log = logger.Logger

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/api/auth/login", controllers.LoginHandler)
	r.GET("/healthz", controllers.HealthCheckHandler)
	r.GET("/boom", controllers.BoomHandler)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("launch server fail:", zap.String("err", err.Error()))
	}
}
