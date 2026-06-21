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

// RegisterRoutes registers the routing endpoints for the user module.
func (m *UserModule) RegisterRoutes(rg *gin.RouterGroup) {
	// Group routes under the version group passed in (e.g. /v1)
	// Matching the old endpoint routes
	rg.GET("/aplikasi/user", m.Handler.DataAll)
	rg.GET("/aplikasi/data", m.Handler.DataAll)
}
