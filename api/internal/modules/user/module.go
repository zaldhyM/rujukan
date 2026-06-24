package user

import (
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

// RegisterAuthRoutes mendaftarkan endpoint publik dengan rate limiting pada login.
func (m *UserModule) RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/auth/login", middleware.RateLimit(m.rdb, 5, time.Minute), m.Handler.Login)
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
