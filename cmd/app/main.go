package main

import (
	"go.uber.org/zap"
	"vincent.com/todo/internal/adapter/http/rest"
	"vincent.com/todo/pkg/auth"
	"vincent.com/todo/pkg/logger"
	"vincent.com/todo/pkg/tracing"
)

var log = logger.Logger

func main() {
	// gin.SetMode(gin.ReleaseMode)
	tracing.Init("todo-backend-service")
	// defer tracing.Closer.Close()
	auth.NewAuthClient(tracing.Tracer)
	// defer authConn.Close()
	r := rest.InitServer()
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("launch server fail:", zap.String("err", err.Error()))
	}
}
