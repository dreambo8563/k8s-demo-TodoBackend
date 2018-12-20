package main

import (
	"vincent.com/todo/internal/adapter/http/rest"
	"vincent.com/todo/internal/pkg/auth"
	"vincent.com/todo/internal/pkg/logger"
	"vincent.com/todo/internal/pkg/tracing"
)

var log = logger.Logger()

func main() {
	// gin.SetMode(gin.ReleaseMode)
	traceClient := tracing.NewTraceClient()
	defer traceClient.Closer.Close()
	authClient := auth.NewAuthClient(traceClient.GetTracer())
	defer authClient.Conn.Close()
	defer log.Sync()
	r := rest.InitServer()
	err := r.Run() // listen and serve on 0.0.0.0:8080
	if err != nil {
		log.Fatal("launch server fail:", log.String("err", err.Error()))
	}
}
