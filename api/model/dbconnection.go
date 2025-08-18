package model

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

func DBConnection() {
	// Load file .env
	err := godotenv.Load("../conf.env")
	if err != nil {
		log.Fatal("Error loading conf.env file")
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
		//defer sqlDB.Close()

		err = sqlDB.Ping()
		if err != nil {
			log.Fatal("❌ Cannot ping database: ", err)
		}
		log.Println("✅ Database ping successful!")
	})
}

func CloseConnectionDB() {
	sqlDB, _ := DB.DB()
	defer sqlDB.Close()
}

func SwitchDatabase(dbName string) error {
	if DB == nil {
		return fmt.Errorf("DB belum diinisialisasi")
	}
	return DB.Exec("USE " + dbName).Error
}
