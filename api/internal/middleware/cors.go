package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware mengizinkan request cross-origin dari frontend rujukan-app.
// Origin yang diizinkan dikonfigurasi via env CORS_ORIGIN (koma-separated).
// Default: http://localhost:3000,http://localhost:5173
func CORSMiddleware() gin.HandlerFunc {
	rawOrigins := os.Getenv("CORS_ORIGIN")
	if rawOrigins == "" {
		rawOrigins = "http://localhost:3000,http://localhost:5173"
	}

	allowed := make(map[string]struct{})
	for _, o := range strings.Split(rawOrigins, ",") {
		allowed[strings.TrimSpace(o)] = struct{}{}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")

		if _, ok := allowed[origin]; ok {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
		}

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
