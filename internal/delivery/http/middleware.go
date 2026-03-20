package http

import (
	"net/http"
	"time"
)

func RateLimitMiddleware(cache interface {
	RateLimit(string, int, time.Duration) (bool, error)
}) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr

			allowed, _ := cache.RateLimit(ip, 10, time.Minute)

			if !allowed {
				http.Error(w, "To Many Request", http.StatusTooManyRequests)
				return
			}

			next(w, r)
		}
	}
}
