package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"vincent.com/todo/internal/adapter/service"
	"vincent.com/todo/internal/pkg/res"
	"vincent.com/todo/internal/pkg/tracing"
)

//UUIDHandler -
func UUIDHandler(c *gin.Context) {
	client := tracing.NewTraceClient()
	span := client.StartSpan("UUIDHandler")

	type uuidRes struct {
		UUID string `json:"uuid"`
	}
	defer span.Finish()
	ctx := tracing.ContextWithSpan(context.Background(), span)
	uuidService := service.InitializeUUIDCase()
	uuid := uuidService.New(ctx)

	res.JSON(c, &uuidRes{
		UUID: uuid,
	})
}
