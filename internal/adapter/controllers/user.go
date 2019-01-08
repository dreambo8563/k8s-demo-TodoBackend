package controllers

import (
	"github.com/gin-gonic/gin"
	"vincent.com/todo/internal/domain/usecase"
	"vincent.com/todo/internal/pkg/res"
)

//UserInfo - get UserInfo
func UserInfo(c *gin.Context) {
	user := c.MustGet("user").(*usecase.User)
	res.JSON(c, user)
}
