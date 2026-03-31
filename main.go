package main

import (
	"fmt"
	"go-scheduler/internal/handlers"
	"go-scheduler/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	apiKey := os.Getenv("RESEND_CLIENT")
	fromEmail := os.Getenv("RESEND_FROM")
	emailService := services.NewEmailService(apiKey, fromEmail)

	apiHandler := handlers.NewApiHandler(emailService)

	mux := http.NewServeMux()
	apiHandler.RegisterHandlers(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		fmt.Printf("Startup error: %v\n", err)
	}
}
