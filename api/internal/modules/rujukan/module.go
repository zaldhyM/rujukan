package rujukan

import (
	"rujukan/internal/middleware"
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

	// Buat rujukan: admin dan faskes
	rg.POST("/rujukan", middleware.RequireRole("admin", "faskes"), m.Handler.Create)

	// Update status: semua role terautentikasi bisa (admin, faskes, dinkeskab, dinkesprov)
	rg.PUT("/rujukan/:id/status", middleware.RequireRole("admin", "faskes", "dinkeskab", "dinkesprov"), m.Handler.UpdateStatus)
}
