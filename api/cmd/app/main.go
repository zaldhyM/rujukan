package main

import (
	"log"
	"rujukan/internal/app"
	"rujukan/internal/infrastructure/cache"
	"rujukan/internal/infrastructure/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, relying on environment variables")
	}

	// Defer closing database and cache connections on main exit
	defer database.Close()
	defer cache.Close()

	// Boot application
	application := app.NewApp()
	application.Start()
}

