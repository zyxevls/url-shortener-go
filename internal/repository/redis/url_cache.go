package redisrepo

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type URLCache struct {
	client *redis.Client
	ctx    context.Context
}

func NewURLCache(r *redis.Client) *URLCache {
	return &URLCache{
		client: r,
		ctx:    context.Background(),
	}
}

func (c *URLCache) Set(code string, original string, ttl time.Duration) error {
	return c.client.Set(c.ctx, "short:"+code, original, ttl).Err()
}

func (c *URLCache) Get(code string) (string, error) {
	return c.client.Get(c.ctx, "short:"+code).Result()
}

func (c *URLCache) IncrementClick(code string) {
	c.client.Incr(c.ctx, "click:"+code)
}
