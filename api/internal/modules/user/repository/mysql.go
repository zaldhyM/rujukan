package repository

import (
	"log"
	"rujukan/internal/infrastructure/database"
	"rujukan/internal/modules/user/domain"

	"gorm.io/gorm"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	QueryAll() ([]domain.User, error)
}

type mysqlRepository struct {
	db *gorm.DB
}

// NewMySQLRepository creates a new UserRepository implementation using GORM.
func NewMySQLRepository(db *gorm.DB) UserRepository {
	return &mysqlRepository{db: db}
}

// QueryAll fetches all users from the user table in the aplikasi database.
func (r *mysqlRepository) QueryAll() ([]domain.User, error) {
	// Switch to aplikasi database to ensure correct schema context
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in repository:", err)
		// We can still try to query, in case the connection is already on the correct database
	}

	var users []domain.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
