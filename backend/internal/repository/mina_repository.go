package repository

import (
	defaultDto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type MinaRepository interface {
	Create(entity *entity.Mina) error
	FindAll(pagination *defaultDto.PaginationRequest) ([]entity.Mina, int64, error)
	FindByID(id uint) (*entity.Mina, error)
	Update(entity *entity.Mina) error
	Delete(id uint) error
}

type minaRepository struct {
	db *gorm.DB
}

func NewMinaRepository(db *gorm.DB) MinaRepository {
	return &minaRepository{db: db}
}

func (r *minaRepository) Create(entity *entity.Mina) error {
	return r.db.Create(entity).Error
}

func (r *minaRepository) FindAll(pagination *defaultDto.PaginationRequest) ([]entity.Mina, int64, error) {
	var entities []entity.Mina
	var total int64

	query := r.db.Model(&entity.Mina{})

	
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

func (r *minaRepository) FindByID(id uint) (*entity.Mina, error) {
	var entity entity.Mina
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *minaRepository) Update(entity *entity.Mina) error {
	return r.db.Save(entity).Error
}

func (r *minaRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Mina{}, id).Error
}
