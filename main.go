package main

import (
	"github.com/gin-gonic/gin"
	"vincent.com/todo/controllers"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.POST("/api/auth/login", controllers.LoginHandler)
	r.GET("/api/auth/health", controllers.HealthCheckHandler)
	r.Run() // listen and serve on 0.0.0.0:8080
}
