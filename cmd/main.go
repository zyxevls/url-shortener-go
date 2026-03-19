package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	handler "github.com/zyxevls/internal/delivery/http"
	"github.com/zyxevls/internal/repository/postgres"
	"github.com/zyxevls/internal/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	repo := postgres.NewURLRepository()
	useCase := usecase.NewUrlUseCase(repo)
	h := handler.NewHandler(useCase)

	http.HandleFunc("/api/v1/shorten", h.CreateShortURL)
	http.HandleFunc("/", h.Redirect)

	log.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}
