package pasien

import (
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
	return &PasienModule{
		Handler: handler,
	}
}

func (m *PasienModule) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/pasien", m.Handler.QueryAll)
	rg.GET("/pasien/:id", m.Handler.GetByID)
	rg.POST("/pasien", m.Handler.Create)
	rg.PUT("/pasien/:id", m.Handler.Update)
	rg.DELETE("/pasien/:id", m.Handler.Delete)
}
