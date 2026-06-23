package repository

import (
	"crypto/rand"
	"fmt"
	"time"

	"rujukan/internal/modules/rujukan/domain"

	"gorm.io/gorm"
)

type RujukanRepository interface {
	QueryAll(search, status string, limit, offset int) ([]domain.Rujukan, int64, error)
	GetByID(id int) (*domain.Rujukan, error)
	Create(rujukan *domain.Rujukan) error
	UpdateStatus(id int, status string) error
}

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) RujukanRepository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) QueryAll(search, status string, limit, offset int) ([]domain.Rujukan, int64, error) {
	var data []domain.Rujukan
	var total int64

	query := r.db.Model(&domain.Rujukan{})

	if search != "" {
		like := "%" + search + "%"
		query = query.Where("NO_RUJUKAN LIKE ? OR ID_PASIEN LIKE ? OR ID_FASKES_ASAL LIKE ? OR ID_FASKES_TUJUAN LIKE ?",
			like, like, like, like)
	}
	if status != "" {
		query = query.Where("STATUS = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Order("TANGGAL DESC").Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, total, nil
}

func (r *mysqlRepository) GetByID(id int) (*domain.Rujukan, error) {
	var data domain.Rujukan
	if err := r.db.First(&data, "ID = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *mysqlRepository) Create(rujukan *domain.Rujukan) error {
	if rujukan.NoRujukan == "" {
		rujukan.NoRujukan = generateNoRujukan()
	}
	return r.db.Create(rujukan).Error
}

func (r *mysqlRepository) UpdateStatus(id int, status string) error {
	return r.db.Model(&domain.Rujukan{}).Where("ID = ?", id).Update("STATUS", status).Error
}

func generateNoRujukan() string {
	b := make([]byte, 4)
	_, _ = rand.Read(b)
	return fmt.Sprintf("RJK-%s-%X", time.Now().Format("20060102"), b)
}
