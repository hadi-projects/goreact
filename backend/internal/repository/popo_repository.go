package repository

import (
	defaultdto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type PopoRepository interface {
	Create(entity *entity.Popo) error
	FindAll(pagination *defaultdto.PaginationRequest) ([]entity.Popo, int64, error)
	FindByID(id uint) (*entity.Popo, error)
	Update(entity *entity.Popo) error
	Delete(id uint) error
}

type popoRepository struct {
	db *gorm.DB
}

func NewPopoRepository(db *gorm.DB) PopoRepository {
	return &popoRepository{db: db}
}

func (r *popoRepository) Create(entity *entity.Popo) error {
	return r.db.Create(entity).Error
}

func (r *popoRepository) FindAll(pagination *defaultdto.PaginationRequest) ([]entity.Popo, int64, error) {
	var entities []entity.Popo
	var total int64

	query := r.db.Model(&entity.Popo{})

	
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

func (r *popoRepository) FindByID(id uint) (*entity.Popo, error) {
	var entity entity.Popo
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *popoRepository) Update(entity *entity.Popo) error {
	return r.db.Save(entity).Error
}

func (r *popoRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Popo{}, id).Error
}
