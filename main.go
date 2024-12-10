package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	r "github.com/tolma-ch/tolmach-go/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := r.MainRouter()
	srv := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
