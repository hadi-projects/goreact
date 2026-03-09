package repository

import (
	defaultdto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type TesttRepository interface {
	Create(entity *entity.Testt) error
	FindAll(pagination *defaultdto.PaginationRequest) ([]entity.Testt, int64, error)
	FindByID(id uint) (*entity.Testt, error)
	Update(entity *entity.Testt) error
	Delete(id uint) error
}

type testtRepository struct {
	db *gorm.DB
}

func NewTesttRepository(db *gorm.DB) TesttRepository {
	return &testtRepository{db: db}
}

func (r *testtRepository) Create(entity *entity.Testt) error {
	return r.db.Create(entity).Error
}

func (r *testtRepository) FindAll(pagination *defaultdto.PaginationRequest) ([]entity.Testt, int64, error) {
	var entities []entity.Testt
	var total int64

	query := r.db.Model(&entity.Testt{})

	
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

func (r *testtRepository) FindByID(id uint) (*entity.Testt, error) {
	var entity entity.Testt
	err := r.db.First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *testtRepository) Update(entity *entity.Testt) error {
	return r.db.Save(entity).Error
}

func (r *testtRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Testt{}, id).Error
}
