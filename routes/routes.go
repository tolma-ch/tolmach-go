package routes

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	cont "github.com/tolma-ch/tolmach-go/controller"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

func MainRouter(swaggerUIdir embed.FS) (*gin.Engine, *fizz.Fizz) {
	router := gin.Default()
	f := fizz.NewFromEngine(router)
	f.GET("/task/:id", nil, tonic.Handler(cont.GetTask, http.StatusOK))

	info := &openapi.Info{
		Title:       "Task API",
		Description: "manage tasks",
		Version:     "0.0.1",
	}
	f.GET("/openapi.json", nil, f.OpenAPI(info, "json"))

	swaggerAssets, fsErr := fs.Sub(swaggerUIdir, "swagger-ui")
	if fsErr != nil {
		log.Fatalf("Failed to load embedded Swagger UI assets: %v", fsErr)
	}
	router.StaticFS("/docs", http.FS(swaggerAssets))

	return router, f
}
