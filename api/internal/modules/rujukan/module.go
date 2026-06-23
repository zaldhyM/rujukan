package rujukan

import (
	"rujukan/internal/modules/rujukan/delivery/http"
	"rujukan/internal/modules/rujukan/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RujukanModule struct {
	Handler *http.RujukanHandler
}

func NewRujukanModule(db *gorm.DB) *RujukanModule {
	repo := repository.NewMySQLRepository(db)
	handler := http.NewRujukanHandler(repo)
	return &RujukanModule{Handler: handler}
}

func (m *RujukanModule) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/rujukan", m.Handler.QueryAll)
	rg.GET("/rujukan/:id", m.Handler.GetByID)
	rg.POST("/rujukan", m.Handler.Create)
	rg.PUT("/rujukan/:id/status", m.Handler.UpdateStatus)
}
