package referensi

import (
	"rujukan/internal/modules/referensi/delivery/http"
	"rujukan/internal/modules/referensi/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReferensiModule struct {
	Handler *http.ReferensiHandler
}

func NewReferensiModule(db *gorm.DB) *ReferensiModule {
	repo := repository.NewMySQLRepository(db)
	handler := http.NewReferensiHandler(repo)
	return &ReferensiModule{
		Handler: handler,
	}
}

func (m *ReferensiModule) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/referensi", m.Handler.QueryReferensi)
	rg.GET("/referensi/:jenis/:id", m.Handler.GetReferensiByKeys)
	rg.GET("/jenis-referensi", m.Handler.QueryJenisReferensi)
	rg.GET("/jenis-referensi/:id", m.Handler.GetJenisReferensiByID)
}
