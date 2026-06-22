package wilayah

import (
	"rujukan/internal/modules/wilayah/delivery/http"
	"rujukan/internal/modules/wilayah/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WilayahModule struct {
	Handler *http.WilayahHandler
}

func NewWilayahModule(db *gorm.DB) *WilayahModule {
	repo := repository.NewMySQLRepository(db)
	handler := http.NewWilayahHandler(repo)
	return &WilayahModule{
		Handler: handler,
	}
}

func (m *WilayahModule) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/wilayah", m.Handler.QueryWilayah)
	rg.GET("/wilayah/:id", m.Handler.GetWilayahByID)
	rg.GET("/jenis-wilayah", m.Handler.QueryJenisWilayah)
	rg.GET("/jenis-wilayah/:id", m.Handler.GetJenisWilayahByID)
}
