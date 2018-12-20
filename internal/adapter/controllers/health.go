package controllers

import (
	"net/http"

	"vincent.com/todo/internal/pkg/auth"
	"vincent.com/todo/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger()

// HealthCheckHandler - handler health check
func HealthCheckHandler(c *gin.Context) {
	err := auth.HealthZ()
	if err != nil {
		log.Error("HealthCheckHandler", log.String("err", err.Error()))
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "ok")
}

// BoomHandler - Boom!
func BoomHandler(c *gin.Context) {
	log.Panic("BoomHandler")
}
