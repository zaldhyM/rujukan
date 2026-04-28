package main

import (
	"log"
	"os"
	v1 "rujukan/internal/delivery/v1"
	"rujukan/model"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	model.DBConnection()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	router := v1.SetupRouter()
	router.Run(":" + port)
}
