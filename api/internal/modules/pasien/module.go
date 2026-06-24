package pasien

import (
	"rujukan/internal/middleware"
	"rujukan/internal/modules/pasien/delivery/http"
	"rujukan/internal/modules/pasien/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PasienModule struct {
	Handler *http.PasienHandler
}

func NewPasienModule(db *gorm.DB) *PasienModule {
	repo := repository.NewMySQLRepository(db)
	handler := http.NewPasienHandler(repo)
	return &PasienModule{Handler: handler}
}

func (m *PasienModule) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/pasien", m.Handler.QueryAll)
	rg.GET("/pasien/:id", m.Handler.GetByID)

	write := rg.Group("", middleware.RequireRole("admin", "faskes"))
	write.POST("/pasien", m.Handler.Create)
	write.PUT("/pasien/:id", m.Handler.Update)

	admin := rg.Group("", middleware.RequireRole("admin"))
	admin.DELETE("/pasien/:id", m.Handler.Delete)
}
