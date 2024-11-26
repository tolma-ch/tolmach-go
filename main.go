package main

import (
	"embed"
	"log"
	"net/http"

	r "github.com/tolma-ch/tolmach-go/routes"
)

//go:embed swagger-ui
var swaggerUIdir embed.FS

func main() {
	router, f := r.MainRouter(swaggerUIdir)

	srv := &http.Server{
		Addr:    "localhost:8000",
		Handler: f,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	router.Run("localhost:8000")
}
