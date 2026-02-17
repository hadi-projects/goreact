package service

import (
	"fmt"
	"math"
	"time"

	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"github.com/hadi-projects/go-react-starter/internal/repository"
	"github.com/hadi-projects/go-react-starter/pkg/cache"
	"github.com/hadi-projects/go-react-starter/pkg/logger"
)

type XyzService interface {
	Create(req dto.CreateXyzRequest) (*dto.XyzResponse, error)
	GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	GetByID(id uint) (*dto.XyzResponse, error)
	Update(id uint, req dto.UpdateXyzRequest) (*dto.XyzResponse, error)
	Delete(id uint) error
}

type xyzService struct {
	repo  repository.XyzRepository
	cache cache.CacheService
}

func NewXyzService(repo repository.XyzRepository, cache cache.CacheService) XyzService {
	return &xyzService{
		repo:  repo,
		cache: cache,
	}
}

func (s *xyzService) Create(req dto.CreateXyzRequest) (*dto.XyzResponse, error) {
	entity := &entity.Xyz{
		Name: req.Name,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("xyz:*")

	
	logger.AuditLogger.Info().
		Uint("xyz_id", entity.ID).
		Str("action", "xyz_creation").
		Msg("xyz created")
	

	return s.mapToResponse(entity), nil
}

func (s *xyzService) GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("xyz:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached dto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.XyzResponse
	for _, e := range entities {
		responses = append(responses, *s.mapToResponse(&e))
	}

	response := &dto.PaginationResponse{
		Data: responses,
		Meta: dto.PaginationMeta{
			CurrentPage: pagination.GetPage(),
			TotalPages:  int(math.Ceil(float64(total) / float64(pagination.GetLimit()))),
			TotalData:   total,
			Limit:       pagination.GetLimit(),
		},
	}

	s.cache.Set(cacheKey, response, 5*time.Minute)
	return response, nil
}

func (s *xyzService) GetByID(id uint) (*dto.XyzResponse, error) {
	cacheKey := fmt.Sprintf("xyz:%d", id)
	var cached dto.XyzResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entity, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := s.mapToResponse(entity)
	s.cache.Set(cacheKey, response, 5*time.Minute)
	return response, nil
}

func (s *xyzService) Update(id uint, req dto.UpdateXyzRequest) (*dto.XyzResponse, error) {
	entity, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if req.Name != "" {
		entity.Name = req.Name
	}

	if err := s.repo.Update(entity); err != nil {
		return nil, err
	}

	s.cache.Delete(fmt.Sprintf("xyz:%d", id))
	s.cache.DeletePattern("xyz:*")

	
	logger.AuditLogger.Info().
		Uint("xyz_id", entity.ID).
		Str("action", "xyz_update").
		Msg("xyz updated")
	

	return s.mapToResponse(entity), nil
}

func (s *xyzService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("xyz:%d", id))
	s.cache.DeletePattern("xyz:*")

	
	logger.AuditLogger.Info().
		Uint("xyz_id", id).
		Str("action", "xyz_deletion").
		Msg("xyz deleted")
	

	return s.repo.Delete(id)
}

func (s *xyzService) mapToResponse(entity *entity.Xyz) *dto.XyzResponse {
	return &dto.XyzResponse{
		ID:        entity.ID,
		Name: entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
