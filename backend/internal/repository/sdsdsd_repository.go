package repository

import (
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type SdsdsdRepository interface {
	Create(entity *entity.Sdsdsd) error
	FindAll(pagination *dto.PaginationRequest) ([]entity.Sdsdsd, int64, error)
	FindByID(id uint) (*entity.Sdsdsd, error)
	Update(entity *entity.Sdsdsd) error
	Delete(id uint) error
}

type sdsdsdRepository struct {
	db *gorm.DB
}

func NewSdsdsdRepository(db *gorm.DB) SdsdsdRepository {
	return &sdsdsdRepository{db: db}
}

func (r *sdsdsdRepository) Create(entity *entity.Sdsdsd) error {
	return r.db.Create(entity).Error
}

func (r *sdsdsdRepository) FindAll(pagination *dto.PaginationRequest) ([]entity.Sdsdsd, int64, error) {
	var entities []entity.Sdsdsd
	var total int64

	query := r.db.Model(&entity.Sdsdsd{})

	
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

func (r *sdsdsdRepository) FindByID(id uint) (*entity.Sdsdsd, error) {
	var entity entity.Sdsdsd
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *sdsdsdRepository) Update(entity *entity.Sdsdsd) error {
	return r.db.Save(entity).Error
}

func (r *sdsdsdRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Sdsdsd{}, id).Error
}
