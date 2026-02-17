package repository

import (
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type AkusajaRepository interface {
	Create(entity *entity.Akusaja) error
	FindAll(pagination *dto.PaginationRequest) ([]entity.Akusaja, int64, error)
	FindByID(id uint) (*entity.Akusaja, error)
	Update(entity *entity.Akusaja) error
	Delete(id uint) error
}

type akusajaRepository struct {
	db *gorm.DB
}

func NewAkusajaRepository(db *gorm.DB) AkusajaRepository {
	return &akusajaRepository{db: db}
}

func (r *akusajaRepository) Create(entity *entity.Akusaja) error {
	return r.db.Create(entity).Error
}

func (r *akusajaRepository) FindAll(pagination *dto.PaginationRequest) ([]entity.Akusaja, int64, error) {
	var entities []entity.Akusaja
	var total int64

	query := r.db.Model(&entity.Akusaja{})

	
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

func (r *akusajaRepository) FindByID(id uint) (*entity.Akusaja, error) {
	var entity entity.Akusaja
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *akusajaRepository) Update(entity *entity.Akusaja) error {
	return r.db.Save(entity).Error
}

func (r *akusajaRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Akusaja{}, id).Error
}
