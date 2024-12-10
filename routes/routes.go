package routes

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	cont "github.com/tolma-ch/tolmach-go/controllers"
	"github.com/tolma-ch/tolmach-go/controllers/user"
	"github.com/tolma-ch/tolmach-go/db"
	"github.com/tolma-ch/tolmach-go/docs"
	mw "github.com/tolma-ch/tolmach-go/middleware"
)

func MainRouter() *gin.Engine {
	d := db.Init()
	d.AutoMigrate(&user.User{})

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/task/:id", mw.AuthMiddleware(), cont.GetTask)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user.UserRouter(r)

	return r
}
