package controllers

import (
	"net/http"

	"go.uber.org/zap"
	"vincent.com/todo/service/auth"
	"vincent.com/todo/service/logger"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger

// HealthCheckHandler - handler health check
func HealthCheckHandler(c *gin.Context) {

	err := auth.HealthZ()
	if err != nil {
		log.Error("HealthCheckHandler", zap.String("err", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}
