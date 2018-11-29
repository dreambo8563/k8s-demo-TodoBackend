package controllers

import (
	"go.uber.org/zap"
	"vincent.com/todo/service/logger"
	"vincent.com/todo/service/res"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger

// HealthCheckHandler - handler health check
func HealthCheckHandler(c *gin.Context) {
	log.Info("health check", zap.String("status", "ok"))
	res.JSON(c, "ok")
}
