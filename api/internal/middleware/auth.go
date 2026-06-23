package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"rujukan/internal/infrastructure/auth"
	"rujukan/internal/infrastructure/cache"
	"rujukan/internal/infrastructure/database"
	"rujukan/internal/modules/user/domain"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// AuthMiddleware authenticates requests using JWT (Bearer token or session cookie).
func AuthMiddleware(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenStr string

		// 1. Try to get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
				tokenStr = parts[1]
			}
		}

		// 2. If not found in header, try to get from session_token cookie
		if tokenStr == "" {
			if cookie, err := c.Cookie("session_token"); err == nil {
				tokenStr = cookie
			}
		}

		// 3. If still empty, deny access
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authentication required. Provide a Bearer token or session cookie.",
			})
			c.Abort()
			return
		}

		// 4. Validate JWT signature and expiry
		claims, err := auth.ValidateToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Invalid or expired token: " + err.Error(),
			})
			c.Abort()
			return
		}

		// 5. Verify active session against Redis (single active session enforcement)
		sessionToken, err := rdb.Get(context.Background(), cache.SessionKey(claims.UserID)).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Session has been invalidated or logged out",
			})
			c.Abort()
			return
		}
		if sessionToken != tokenStr {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Session has been invalidated or logged out",
			})
			c.Abort()
			return
		}

		// 6. Verify user status and masa aktif from database
		if err := database.Switch("aplikasi"); err != nil {
			log.Println("⚠️ Error switching to database aplikasi in auth middleware:", err)
		}

		var dbUser domain.User
		if err := database.DB.First(&dbUser, "ID = ?", claims.UserID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "User not found or database error",
			})
			c.Abort()
			return
		}

		if dbUser.STATUS != 1 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "User account is inactive",
			})
			c.Abort()
			return
		}

		if dbUser.MASA_AKTIF != "" {
			if expiryDate, err := time.Parse("2006-01-02", dbUser.MASA_AKTIF); err == nil {
				if time.Now().After(expiryDate) {
					c.JSON(http.StatusUnauthorized, gin.H{
						"success": false,
						"message": "User account has expired",
					})
					c.Abort()
					return
				}
			}
		}

		// Store user info in context for use in handlers
		c.Set("userID", dbUser.ID)
		c.Set("username", dbUser.USERNAME)
		c.Set("role", dbUser.ROLE)
		c.Set("user", dbUser)

		c.Next()
	}
}
