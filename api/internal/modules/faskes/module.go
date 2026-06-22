package faskes

import (
	"rujukan/internal/modules/faskes/delivery/http"
	"rujukan/internal/modules/faskes/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FaskesModule struct {
	Handler *http.FaskesHandler
}

func NewFaskesModule(db *gorm.DB) *FaskesModule {
	repo := repository.NewMySQLRepository(db)
	handler := http.NewFaskesHandler(repo)
	return &FaskesModule{
		Handler: handler,
	}
}

func (m *FaskesModule) RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/faskes", m.Handler.QueryAll)
	rg.GET("/faskes/:id", m.Handler.GetByID)
	rg.POST("/faskes", m.Handler.Create)
	rg.PUT("/faskes/:id", m.Handler.Update)
	rg.DELETE("/faskes/:id", m.Handler.Delete)
}
