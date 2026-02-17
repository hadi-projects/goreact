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

type NinaService interface {
	Create(req dto.CreateNinaRequest) (*dto.NinaResponse, error)
	GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	GetByID(id uint) (*dto.NinaResponse, error)
	Update(id uint, req dto.UpdateNinaRequest) (*dto.NinaResponse, error)
	Delete(id uint) error
}

type ninaService struct {
	repo  repository.NinaRepository
	cache cache.CacheService
}

func NewNinaService(repo repository.NinaRepository, cache cache.CacheService) NinaService {
	return &ninaService{
		repo:  repo,
		cache: cache,
	}
}

func (s *ninaService) Create(req dto.CreateNinaRequest) (*dto.NinaResponse, error) {
	entity := &entity.Nina{
		Names: req.Names,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("nina:*")

	
	logger.AuditLogger.Info().
		Uint("nina_id", entity.ID).
		Str("action", "nina_creation").
		Msg("nina created")
	

	return s.mapToResponse(entity), nil
}

func (s *ninaService) GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("nina:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached dto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.NinaResponse
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

func (s *ninaService) GetByID(id uint) (*dto.NinaResponse, error) {
	cacheKey := fmt.Sprintf("nina:%d", id)
	var cached dto.NinaResponse
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

func (s *ninaService) Update(id uint, req dto.UpdateNinaRequest) (*dto.NinaResponse, error) {
	entity, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if req.Names != "" {
		entity.Names = req.Names
	}

	if err := s.repo.Update(entity); err != nil {
		return nil, err
	}

	s.cache.Delete(fmt.Sprintf("nina:%d", id))
	s.cache.DeletePattern("nina:*")

	
	logger.AuditLogger.Info().
		Uint("nina_id", entity.ID).
		Str("action", "nina_update").
		Msg("nina updated")
	

	return s.mapToResponse(entity), nil
}

func (s *ninaService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("nina:%d", id))
	s.cache.DeletePattern("nina:*")

	
	logger.AuditLogger.Info().
		Uint("nina_id", id).
		Str("action", "nina_deletion").
		Msg("nina deleted")
	

	return s.repo.Delete(id)
}

func (s *ninaService) mapToResponse(entity *entity.Nina) *dto.NinaResponse {
	return &dto.NinaResponse{
		ID:        entity.ID,
		Names: entity.Names,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
