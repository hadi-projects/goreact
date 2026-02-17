package repository

import (
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type XyzRepository interface {
	Create(entity *entity.Xyz) error
	FindAll(pagination *dto.PaginationRequest) ([]entity.Xyz, int64, error)
	FindByID(id uint) (*entity.Xyz, error)
	Update(entity *entity.Xyz) error
	Delete(id uint) error
}

type xyzRepository struct {
	db *gorm.DB
}

func NewXyzRepository(db *gorm.DB) XyzRepository {
	return &xyzRepository{db: db}
}

func (r *xyzRepository) Create(entity *entity.Xyz) error {
	return r.db.Create(entity).Error
}

func (r *xyzRepository) FindAll(pagination *dto.PaginationRequest) ([]entity.Xyz, int64, error) {
	var entities []entity.Xyz
	var total int64

	query := r.db.Model(&entity.Xyz{})

	
	if pagination.Search != "" {
		query = query.Where("name LIKE ?", "%"+pagination.Search+"%")
	}
	

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (pagination.GetPage() - 1) * pagination.GetLimit()
	err := query.Order("id DESC").
		Limit(pagination.GetLimit()).
		Offset(offset).
		Find(&entities).Error

	return entities, total, err
}

func (r *xyzRepository) FindByID(id uint) (*entity.Xyz, error) {
	var entity entity.Xyz
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *xyzRepository) Update(entity *entity.Xyz) error {
	return r.db.Save(entity).Error
}

func (r *xyzRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Xyz{}, id).Error
}
