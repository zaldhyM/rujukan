package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// Connect initializes the database connection pool using configurations from environment variables.
func Connect() *gorm.DB {
	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, relying on environment variables")
	}

	// Ambil value
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	once.Do(func() {
		// Format DSN untuk MySQL
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, name)

		// Connect DB pakai GORM
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database: ", err)
		}
		DB = db
		log.Println("✅ Database connected successfully!")

		// Test query sederhana
		sqlDB, _ := db.DB()
		err = sqlDB.Ping()
		if err != nil {
			log.Fatal("❌ Cannot ping database: ", err)
		}
		log.Println("✅ Database ping successful!")
	})

	return DB
}

// Close closes the database connection.
func Close() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Println("Error getting database instance to close:", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Println("Error closing database connection:", err)
		} else {
			log.Println("✅ Database connection closed successfully!")
		}
	}
}

// Switch switches the active database schema for the current session.
func Switch(dbName string) error {
	if DB == nil {
		return fmt.Errorf("DB belum diinisialisasi")
	}
	return DB.Exec("USE " + dbName).Error
}
