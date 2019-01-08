package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response struct
// type Response struct {
// 	code int
// 	data interface{}
// 	msg  string
// }

// JSON - send json response
func JSON(g *gin.Context, data interface{}) {
	g.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

// Err - send err to frontend
func Err(g *gin.Context, msg string) {
	g.JSON(http.StatusOK, gin.H{"success": false, "data": "", "msg": msg})
}

//Abort -
func Abort(g *gin.Context, code int) {
	g.AbortWithStatus(code)
}
