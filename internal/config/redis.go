package config

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedisClient() *redis.Client {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	_, err := rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Redis connection error", err)
	}

	log.Println("Connected to redis⚡")
	return rdb
}
