package repository

import (
	"crypto/rand"
	"fmt"
	"rujukan/internal/modules/pasien/domain"
	"gorm.io/gorm"
)

type PasienRepository interface {
	QueryAll(search string, limit, offset int) ([]domain.Pasien, int64, error)
	GetByID(id string) (*domain.Pasien, error)
	Create(pasien *domain.Pasien) error
	Update(pasien *domain.Pasien) error
	Delete(id string) error
}

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) PasienRepository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) QueryAll(search string, limit, offset int) ([]domain.Pasien, int64, error) {
	var pasien []domain.Pasien
	var total int64

	query := r.db.Model(&domain.Pasien{})
	if search != "" {
		query = query.Where("NAMA LIKE ? OR NIK LIKE ? OR NO_JKN LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Find(&pasien).Error; err != nil {
		return nil, 0, err
	}

	return pasien, total, nil
}

func (r *mysqlRepository) GetByID(id string) (*domain.Pasien, error) {
	var pasien domain.Pasien
	if err := r.db.First(&pasien, "ID = ?", id).Error; err != nil {
		return nil, err
	}
	return &pasien, nil
}

func (r *mysqlRepository) Create(pasien *domain.Pasien) error {
	if pasien.ID == "" {
		pasien.ID = generateUUID()
	}
	return r.db.Create(pasien).Error
}

func (r *mysqlRepository) Update(pasien *domain.Pasien) error {
	return r.db.Save(pasien).Error
}

func (r *mysqlRepository) Delete(id string) error {
	return r.db.Delete(&domain.Pasien{}, "ID = ?", id).Error
}

func generateUUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
