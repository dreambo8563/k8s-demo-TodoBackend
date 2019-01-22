package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"vincent.com/todo/internal/pkg/uc"

	"vincent.com/todo/internal/adapter/http/rest"
	"vincent.com/todo/internal/pkg/auth"
	"vincent.com/todo/internal/pkg/logger"
	"vincent.com/todo/internal/pkg/tracing"
)

var log = logger.Logger()

func main() {
	// gin.SetMode(gin.ReleaseMode)
	traceClient := tracing.NewTraceClient()
	authClient := auth.NewAuthClient(traceClient.GetTracer())
	r := rest.InitServer()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("launch server fail:", log.String("err", err.Error()))
		}
	}()

	// err := r.Run() // listen and serve on 0.0.0.0:8080
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tracing.Close()
	authClient.Close()
	err := log.Sync()
	uc.Destroy()
	if err != nil {
		log.Error("log.Sync error")
	}
	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", log.String("err", err.Error()))
	}
	log.Info("Server exiting")

}
