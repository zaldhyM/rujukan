package repository

import (
	"rujukan/internal/modules/wilayah/domain"
	"gorm.io/gorm"
)

type WilayahRepository interface {
	QueryWilayah(search string, jenis int8, parentID string, limit, offset int) ([]domain.Wilayah, int64, error)
	GetWilayahByID(id string) (*domain.Wilayah, error)
	QueryJenisWilayah() ([]domain.JenisWilayah, error)
	GetJenisWilayahByID(id int8) (*domain.JenisWilayah, error)
}

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) WilayahRepository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) QueryWilayah(search string, jenis int8, parentID string, limit, offset int) ([]domain.Wilayah, int64, error) {
	var wilayah []domain.Wilayah
	var total int64

	query := r.db.Model(&domain.Wilayah{})

	if search != "" {
		query = query.Where("DESKRIPSI LIKE ? OR ID LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if jenis > 0 {
		query = query.Where("JENIS = ?", jenis)
	}

	if parentID != "" {
		// Filter by parent ID. For example: if parent is '11' (Provinsi Aceh), 
		// its Kabupaten will have IDs starting with '11' (e.g. '1101').
		query = query.Where("ID LIKE ?", parentID+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Find(&wilayah).Error; err != nil {
		return nil, 0, err
	}

	return wilayah, total, nil
}

func (r *mysqlRepository) GetWilayahByID(id string) (*domain.Wilayah, error) {
	var wilayah domain.Wilayah
	if err := r.db.First(&wilayah, "ID = ?", id).Error; err != nil {
		return nil, err
	}
	return &wilayah, nil
}

func (r *mysqlRepository) QueryJenisWilayah() ([]domain.JenisWilayah, error) {
	var jenis []domain.JenisWilayah
	if err := r.db.Find(&jenis).Error; err != nil {
		return nil, err
	}
	return jenis, nil
}

func (r *mysqlRepository) GetJenisWilayahByID(id int8) (*domain.JenisWilayah, error) {
	var jenis domain.JenisWilayah
	if err := r.db.First(&jenis, id).Error; err != nil {
		return nil, err
	}
	return &jenis, nil
}
