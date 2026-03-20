package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/zyxevls/internal/config"
	handler "github.com/zyxevls/internal/delivery/http"
	"github.com/zyxevls/internal/repository/postgres"
	redisrepo "github.com/zyxevls/internal/repository/redis"
	"github.com/zyxevls/internal/usecase"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	db := config.NewPostgresDB()
	rdb := config.NewRedisClient()

	repo := postgres.NewURLRepository(db)
	cache := redisrepo.NewURLCache(rdb)

	usecase := usecase.NewUrlUseCase(repo, cache)
	h := handler.NewHandler(usecase)

	http.HandleFunc("/api/v1/shorten", h.CreateShortURL)
	http.HandleFunc("/", h.Redirect)

	log.Println("Server is running on port :8080")
	http.ListenAndServe(":8080", nil)
}
