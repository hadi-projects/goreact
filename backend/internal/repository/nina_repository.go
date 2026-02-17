package repository

import (
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type NinaRepository interface {
	Create(entity *entity.Nina) error
	FindAll(pagination *dto.PaginationRequest) ([]entity.Nina, int64, error)
	FindByID(id uint) (*entity.Nina, error)
	Update(entity *entity.Nina) error
	Delete(id uint) error
}

type ninaRepository struct {
	db *gorm.DB
}

func NewNinaRepository(db *gorm.DB) NinaRepository {
	return &ninaRepository{db: db}
}

func (r *ninaRepository) Create(entity *entity.Nina) error {
	return r.db.Create(entity).Error
}

func (r *ninaRepository) FindAll(pagination *dto.PaginationRequest) ([]entity.Nina, int64, error) {
	var entities []entity.Nina
	var total int64

	query := r.db.Model(&entity.Nina{})

	
	if pagination.Search != "" {
		query = query.Where("names LIKE ?", "%"+pagination.Search+"%")
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

func (r *ninaRepository) FindByID(id uint) (*entity.Nina, error) {
	var entity entity.Nina
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *ninaRepository) Update(entity *entity.Nina) error {
	return r.db.Save(entity).Error
}

func (r *ninaRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Nina{}, id).Error
}
