package routes

import (
	"github.com/gin-gonic/gin"
	cont "github.com/tolma-ch/tolmach-go/controllers"
	"github.com/tolma-ch/tolmach-go/controllers/user"
	mw "github.com/tolma-ch/tolmach-go/middleware"
)

func MainRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/task/:id", mw.AuthMiddleware(), cont.GetTask)

	r.POST("/login", user.Login)

	return r
}
