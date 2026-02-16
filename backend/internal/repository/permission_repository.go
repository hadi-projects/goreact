package repository

import (
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"gorm.io/gorm"
)

type PermissionRepository interface {
	Create(permission *entity.Permission) error
	FindAll() ([]entity.Permission, error)
	FindByID(id uint) (*entity.Permission, error)
	Update(permission *entity.Permission) error
	Delete(id uint) error
}

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) Create(permission *entity.Permission) error {
	return r.db.Create(permission).Error
}

func (r *permissionRepository) FindAll() ([]entity.Permission, error) {
	var permissions []entity.Permission
	err := r.db.Find(&permissions).Error
	return permissions, err
}

func (r *permissionRepository) FindByID(id uint) (*entity.Permission, error) {
	var permission entity.Permission
	err := r.db.First(&permission, id).Error
	if err != nil {
		return nil, err
	}
	return &permission, nil
}

func (r *permissionRepository) Update(permission *entity.Permission) error {
	return r.db.Save(permission).Error
}

func (r *permissionRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Permission{}, id).Error
}
