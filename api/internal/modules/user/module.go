package user

import (
	"log"
	"os"
	"strconv"
	"time"

	"rujukan/internal/middleware"
	"rujukan/internal/modules/user/delivery/http"
	"rujukan/internal/modules/user/repository"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserModule struct {
	Handler *http.UserHandler
	rdb     *redis.Client
}

func NewUserModule(db *gorm.DB, rdb *redis.Client) *UserModule {
	repo := repository.NewMySQLRepository(db)
	handler := http.NewUserHandler(repo, rdb)
	return &UserModule{Handler: handler, rdb: rdb}
}

// RegisterAuthRoutes mendaftarkan endpoint publik.
// Rate limiting aktif hanya jika RATE_LIMIT_MAX di-set di .env.
func (m *UserModule) RegisterAuthRoutes(rg *gin.RouterGroup) {
	handlers := []gin.HandlerFunc{}

	if rl := rateLimitFromEnv(m.rdb); rl != nil {
		handlers = append(handlers, rl)
	}

	handlers = append(handlers, m.Handler.Login)
	rg.POST("/auth/login", handlers...)
	rg.POST("/auth/register", m.Handler.Register)
}

// RegisterRoutes mendaftarkan endpoint yang memerlukan autentikasi.
func (m *UserModule) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/auth/logout", m.Handler.Logout)
	rg.GET("/auth/me", m.Handler.Me)

	admin := rg.Group("", middleware.RequireRole("admin"))
	admin.GET("/aplikasi/user", m.Handler.DataAll)
	admin.GET("/aplikasi/data", m.Handler.DataAll)
}

// rateLimitFromEnv membaca konfigurasi rate limiter dari env.
// Mengembalikan nil jika RATE_LIMIT_MAX tidak di-set — rate limiter dinonaktifkan.
func rateLimitFromEnv(rdb *redis.Client) gin.HandlerFunc {
	maxStr := os.Getenv("RATE_LIMIT_MAX")
	if maxStr == "" {
		log.Println("ℹ️  RATE_LIMIT_MAX tidak di-set, rate limiter login dinonaktifkan")
		return nil
	}

	max, err := strconv.ParseInt(maxStr, 10, 64)
	if err != nil || max <= 0 {
		log.Printf("⚠️  RATE_LIMIT_MAX tidak valid ('%s'), rate limiter dinonaktifkan\n", maxStr)
		return nil
	}

	window := time.Minute
	if windowStr := os.Getenv("RATE_LIMIT_WINDOW_SECONDS"); windowStr != "" {
		if secs, err := strconv.Atoi(windowStr); err == nil && secs > 0 {
			window = time.Duration(secs) * time.Second
		}
	}

	log.Printf("✅ Rate limiter login aktif: max %d percobaan per %s\n", max, window)
	return middleware.RateLimit(rdb, max, window)
}
