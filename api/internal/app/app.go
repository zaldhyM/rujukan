package app

import (
	"log"
	"os"
	"rujukan/internal/infrastructure/database"
	"rujukan/internal/middleware"
	"rujukan/internal/modules/faskes"
	"rujukan/internal/modules/pasien"
	"rujukan/internal/modules/referensi"
	"rujukan/internal/modules/user"
	"rujukan/internal/modules/wilayah"

	"github.com/gin-gonic/gin"
)

// App defines the application structure container.
type App struct {
	Router *gin.Engine
}

// NewApp initializes dependencies, databases, modules, and routes.
func NewApp() *App {
	// 1. Initialize Database connection
	db := database.Connect()

	// 2. Set GIN_MODE from env (release/debug)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.DebugMode
	}
	gin.SetMode(ginMode)

	// 3. Create the Gin router engine
	router := gin.Default()

	// Only trust proxies from loopback for production security
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// 4. Initialize Monolith Modules
	userModule := user.NewUserModule(db)
	faskesModule := faskes.NewFaskesModule(db)
	pasienModule := pasien.NewPasienModule(db)
	wilayahModule := wilayah.NewWilayahModule(db)
	referensiModule := referensi.NewReferensiModule(db)

	// 5. Group routes based on version
	v1 := router.Group("/v1")
	{
		// Register public/auth routes
		userModule.RegisterAuthRoutes(v1)

		// Register protected routes group using AuthMiddleware
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			userModule.RegisterRoutes(protected)
			faskesModule.RegisterRoutes(protected)
			pasienModule.RegisterRoutes(protected)
			wilayahModule.RegisterRoutes(protected)
			referensiModule.RegisterRoutes(protected)
		}
	}

	return &App{
		Router: router,
	}
}

// Start runs the HTTP server.
func (a *App) Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // default port fallback
	}

	log.Printf("🚀 Server is starting on port %s", port)
	if err := a.Router.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
