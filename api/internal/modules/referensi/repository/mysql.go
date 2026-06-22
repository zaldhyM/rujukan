package repository

import (
	"rujukan/internal/modules/referensi/domain"
	"gorm.io/gorm"
)

type ReferensiRepository interface {
	QueryReferensi(jenis int8, search string) ([]domain.Referensi, error)
	GetReferensiByKeys(jenis int8, id int16) (*domain.Referensi, error)
	QueryJenisReferensi() ([]domain.JenisReferensi, error)
	GetJenisReferensiByID(id int8) (*domain.JenisReferensi, error)
}

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) ReferensiRepository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) QueryReferensi(jenis int8, search string) ([]domain.Referensi, error) {
	var referensi []domain.Referensi
	query := r.db.Model(&domain.Referensi{})

	if jenis > 0 {
		query = query.Where("JENIS = ?", jenis)
	}

	if search != "" {
		query = query.Where("DESKRIPSI LIKE ?", "%"+search+"%")
	}

	if err := query.Find(&referensi).Error; err != nil {
		return nil, err
	}
	return referensi, nil
}

func (r *mysqlRepository) GetReferensiByKeys(jenis int8, id int16) (*domain.Referensi, error) {
	var ref domain.Referensi
	if err := r.db.First(&ref, "JENIS = ? AND ID = ?", jenis, id).Error; err != nil {
		return nil, err
	}
	return &ref, nil
}

func (r *mysqlRepository) QueryJenisReferensi() ([]domain.JenisReferensi, error) {
	var jenis []domain.JenisReferensi
	if err := r.db.Find(&jenis).Error; err != nil {
		return nil, err
	}
	return jenis, nil
}

func (r *mysqlRepository) GetJenisReferensiByID(id int8) (*domain.JenisReferensi, error) {
	var jenis domain.JenisReferensi
	if err := r.db.First(&jenis, id).Error; err != nil {
		return nil, err
	}
	return &jenis, nil
}
