package main

import (
	"log"
	"net/http"

	r "github.com/tolma-ch/tolmach-go/routes"
)

func main() {
	router := r.MainRouter()

	srv := &http.Server{
		Addr:    "localhost:8000",
		Handler: router,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

	router.Run("localhost:8000")
}
