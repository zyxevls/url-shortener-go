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

func (c *URLCache) RateLimit(ip string, limit int, window time.Duration) (bool, error) {
	key := "rate:" + ip

	count, err := c.client.Incr(c.ctx, key).Result()
	if err != nil {
		return false, err
	}

	if count == 1 {
		c.client.Expire(c.ctx, key, window)
	}

	if int(count) > limit {
		return false, nil
	}

	return true, nil
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
