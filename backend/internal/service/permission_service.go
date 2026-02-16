package service

import (
	"math"

	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"github.com/hadi-projects/go-react-starter/internal/repository"
)

type PermissionService interface {
	Create(req dto.CreatePermissionRequest) (*dto.PermissionResponse, error)
	GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
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

func (s *permissionService) GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	permissions, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.PermissionResponse
	for _, perm := range permissions {
		responses = append(responses, dto.PermissionResponse{
			ID:        perm.ID,
			Name:      perm.Name,
			CreatedAt: perm.CreatedAt,
			UpdatedAt: perm.UpdatedAt,
		})
	}

	return &dto.PaginationResponse{
		Data: responses,
		Meta: dto.PaginationMeta{
			CurrentPage: pagination.GetPage(),
			TotalPages:  int(math.Ceil(float64(total) / float64(pagination.GetLimit()))),
			TotalItems:  total,
			Limit:       pagination.GetLimit(),
		},
	}, nil
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
