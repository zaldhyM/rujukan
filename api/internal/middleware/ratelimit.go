package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// RateLimit membatasi jumlah request per IP dalam window waktu tertentu.
// Key Redis: ratelimit:<action>:<IP>
func RateLimit(rdb *redis.Client, max int64, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		key := "ratelimit:login:" + ip

		count, err := rdb.Incr(context.Background(), key).Result()
		if err != nil {
			c.Next()
			return
		}

		if count == 1 {
			rdb.Expire(context.Background(), key, window)
		}

		if count > max {
			ttl, _ := rdb.TTL(context.Background(), key).Result()
			c.Header("Retry-After", ttl.String())
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Terlalu banyak percobaan login. Coba lagi dalam beberapa saat.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
