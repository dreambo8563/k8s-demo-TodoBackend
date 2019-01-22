package middleware

import (
	"fmt"
	"net/http"

	"vincent.com/todo/internal/adapter/service"
	"vincent.com/todo/internal/pkg/logger"

	"vincent.com/todo/internal/pkg/res"

	"github.com/gin-gonic/gin"
)

var log = logger.Logger()

//Auth -
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			res.Abort(c, http.StatusUnauthorized)
			return
		}
		userService := service.InitializeUserCase()
		user, err := userService.GetInfo(c, token)
		fmt.Println(token, err)
		if err != nil {
			res.Abort(c, http.StatusUnauthorized)
			return
		}
		fmt.Println("after Abort")
		log.Sugar().Infof("user %v,", user)
		// Set example variable

		c.Set("user", user)
		c.Next()
	}
}
