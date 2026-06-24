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
	GetByID(id uint) (domain.User, error)
	GetByUsername(username string) (domain.User, error)
	Create(user *domain.User) error
	Update(user *domain.User) error
	UpdateToken(id uint, token string) error
	UpdatePassword(id uint, hashedPassword string) error
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
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in QueryAll:", err)
	}

	var users []domain.User
	cols := "ID, USERNAME, NAMA, NIK, ID_FASKES, MASA_AKTIF, ROLE, STATUS, STATUS_TELEGRAM, ID_TELEGRAM, TERAKHIR_UBAH_PASSWORD"
	if err := r.db.Select(cols).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetByID fetches a user by ID from the aplikasi database.
func (r *mysqlRepository) GetByID(id uint) (domain.User, error) {
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in GetByID:", err)
	}

	var user domain.User
	if err := r.db.First(&user, "ID = ?", id).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// GetByUsername fetches a user by username from the aplikasi database.
func (r *mysqlRepository) GetByUsername(username string) (domain.User, error) {
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in GetByUsername:", err)
	}

	var user domain.User
	if err := r.db.First(&user, "USERNAME = ?", username).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

// Create inserts a new user record into the aplikasi database.
func (r *mysqlRepository) Create(user *domain.User) error {
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in Create:", err)
	}
	return r.db.Create(user).Error
}

// Update updates an existing user record.
func (r *mysqlRepository) Update(user *domain.User) error {
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in Update:", err)
	}
	return r.db.Save(user).Error
}

// UpdateToken updates the session/auth token for a user.
func (r *mysqlRepository) UpdateToken(id uint, token string) error {
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in UpdateToken:", err)
	}
	return r.db.Model(&domain.User{}).Where("ID = ?", id).Update("TOKEN", token).Error
}

// UpdatePassword updates the password hash for a user.
func (r *mysqlRepository) UpdatePassword(id uint, hashedPassword string) error {
	if err := database.Switch("aplikasi"); err != nil {
		log.Println("⚠️ Error switching to database aplikasi in UpdatePassword:", err)
	}
	return r.db.Model(&domain.User{}).Where("ID = ?", id).Update("PASSWORD", hashedPassword).Error
}
