package repository

import (
	"rujukan/internal/modules/faskes/domain"
	"gorm.io/gorm"
)

type FaskesRepository interface {
	QueryAll(search string, limit, offset int) ([]domain.Faskes, int64, error)
	GetByID(id int16) (*domain.Faskes, error)
	Create(faskes *domain.Faskes) error
	Update(faskes *domain.Faskes) error
	Delete(id int16) error
}

type mysqlRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) FaskesRepository {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) QueryAll(search string, limit, offset int) ([]domain.Faskes, int64, error) {
	var faskes []domain.Faskes
	var total int64

	query := r.db.Model(&domain.Faskes{})
	if search != "" {
		query = query.Where("NAMA LIKE ? OR KODE LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	if err := query.Find(&faskes).Error; err != nil {
		return nil, 0, err
	}

	return faskes, total, nil
}

func (r *mysqlRepository) GetByID(id int16) (*domain.Faskes, error) {
	var faskes domain.Faskes
	if err := r.db.First(&faskes, id).Error; err != nil {
		return nil, err
	}
	return &faskes, nil
}

func (r *mysqlRepository) Create(faskes *domain.Faskes) error {
	return r.db.Create(faskes).Error
}

func (r *mysqlRepository) Update(faskes *domain.Faskes) error {
	return r.db.Save(faskes).Error
}

func (r *mysqlRepository) Delete(id int16) error {
	return r.db.Delete(&domain.Faskes{}, id).Error
}
