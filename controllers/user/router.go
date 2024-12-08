package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.Engine {
	r.POST("/login", Login)

	return r
}
