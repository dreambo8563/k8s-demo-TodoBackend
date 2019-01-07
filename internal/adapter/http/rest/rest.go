package rest

import (
	"github.com/gin-gonic/gin"
	"vincent.com/todo/internal/adapter/controllers"
)

//InitServer - init rest server
func InitServer() *gin.Engine {
	r := gin.Default()
	r.POST("/api/auth/login", controllers.RegisterHandler)
	r.GET("/healthz", controllers.HealthCheckHandler)
	r.GET("/user", controllers.UserInfo)
	r.GET("/boom", controllers.BoomHandler)
	r.GET("/uuid", controllers.UUIDHandler)
	return r
}
