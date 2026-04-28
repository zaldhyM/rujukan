package v1

import (
	"os"
	"rujukan/controller/aplikasi/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Set GIN_MODE dari env (release/debug)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	router := gin.Default()

	// Hanya percaya proxy dari loopback (aman untuk production)
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// grouping based on version
	v1 := router.Group("/v1")
	{
		// grouping based on area
		aplikasiRoutes(v1)
		userRoutes(v1)
	}
	return router
}

func userRoutes(rg *gin.RouterGroup) {
	rg.GET("/aplikasi/user", user.DataAll)
}

func aplikasiRoutes(rg *gin.RouterGroup) {
	rg.GET("/aplikasi/data", user.DataAll)
}

