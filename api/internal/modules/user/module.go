package user

import (
	"rujukan/internal/modules/user/delivery/http"
	"rujukan/internal/modules/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserModule represents the user domain module in the monolith.
type UserModule struct {
	Handler *http.UserHandler
}

// NewUserModule initializes and wires up the user module's dependencies.
func NewUserModule(db *gorm.DB) *UserModule {
	repo := repository.NewMySQLRepository(db)
	handler := http.NewUserHandler(repo)
	return &UserModule{
		Handler: handler,
	}
}

// RegisterAuthRoutes registers the authentication endpoints that DO NOT require middleware.
func (m *UserModule) RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/auth/login", m.Handler.Login)
	rg.POST("/auth/register", m.Handler.Register)
}

// RegisterRoutes registers the routing endpoints for the user module that require authentication.
func (m *UserModule) RegisterRoutes(rg *gin.RouterGroup) {
	// Group routes under the version group passed in (e.g. /v1)
	rg.GET("/aplikasi/user", m.Handler.DataAll)
	rg.GET("/aplikasi/data", m.Handler.DataAll)
	rg.POST("/auth/logout", m.Handler.Logout)
	rg.GET("/auth/me", m.Handler.Me)
}
