package repository

import (
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type MakanRepository interface {
	Create(entity *entity.Makan) error
	FindAll(pagination *dto.PaginationRequest) ([]entity.Makan, int64, error)
	FindByID(id uint) (*entity.Makan, error)
	Update(entity *entity.Makan) error
	Delete(id uint) error
}

type makanRepository struct {
	db *gorm.DB
}

func NewMakanRepository(db *gorm.DB) MakanRepository {
	return &makanRepository{db: db}
}

func (r *makanRepository) Create(entity *entity.Makan) error {
	return r.db.Create(entity).Error
}

func (r *makanRepository) FindAll(pagination *dto.PaginationRequest) ([]entity.Makan, int64, error) {
	var entities []entity.Makan
	var total int64

	query := r.db.Model(&entity.Makan{})

	
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

func (r *makanRepository) FindByID(id uint) (*entity.Makan, error) {
	var entity entity.Makan
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *makanRepository) Update(entity *entity.Makan) error {
	return r.db.Save(entity).Error
}

func (r *makanRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Makan{}, id).Error
}
