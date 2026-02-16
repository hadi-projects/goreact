package service

import (
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"github.com/hadi-projects/go-react-starter/internal/repository"
)

type PermissionService interface {
	Create(req dto.CreatePermissionRequest) (*dto.PermissionResponse, error)
	GetAll() ([]dto.PermissionResponse, error)
	Update(id uint, req dto.UpdatePermissionRequest) (*dto.PermissionResponse, error)
	Delete(id uint) error
}

type permissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) PermissionService {
	return &permissionService{repo: repo}
}

func (s *permissionService) Create(req dto.CreatePermissionRequest) (*dto.PermissionResponse, error) {
	permission := &entity.Permission{
		Name: req.Name,
	}

	if err := s.repo.Create(permission); err != nil {
		return nil, err
	}

	return &dto.PermissionResponse{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}, nil
}

func (s *permissionService) GetAll() ([]dto.PermissionResponse, error) {
	permissions, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var response []dto.PermissionResponse
	for _, p := range permissions {
		response = append(response, dto.PermissionResponse{
			ID:        p.ID,
			Name:      p.Name,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return response, nil
}

func (s *permissionService) Update(id uint, req dto.UpdatePermissionRequest) (*dto.PermissionResponse, error) {
	permission, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	permission.Name = req.Name
	if err := s.repo.Update(permission); err != nil {
		return nil, err
	}

	return &dto.PermissionResponse{
		ID:        permission.ID,
		Name:      permission.Name,
		CreatedAt: permission.CreatedAt,
		UpdatedAt: permission.UpdatedAt,
	}, nil
}

func (s *permissionService) Delete(id uint) error {
	return s.repo.Delete(id)
}
