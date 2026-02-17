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

type MakanService interface {
	Create(req dto.CreateMakanRequest) (*dto.MakanResponse, error)
	GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	GetByID(id uint) (*dto.MakanResponse, error)
	Update(id uint, req dto.UpdateMakanRequest) (*dto.MakanResponse, error)
	Delete(id uint) error
}

type makanService struct {
	repo  repository.MakanRepository
	cache cache.CacheService
}

func NewMakanService(repo repository.MakanRepository, cache cache.CacheService) MakanService {
	return &makanService{
		repo:  repo,
		cache: cache,
	}
}

func (s *makanService) Create(req dto.CreateMakanRequest) (*dto.MakanResponse, error) {
	entity := &entity.Makan{
		Name: req.Name,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("makan:*")

	
	logger.AuditLogger.Info().
		Uint("makan_id", entity.ID).
		Str("action", "makan_creation").
		Msg("makan created")
	

	return s.mapToResponse(entity), nil
}

func (s *makanService) GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("makan:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached dto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.MakanResponse
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

func (s *makanService) GetByID(id uint) (*dto.MakanResponse, error) {
	cacheKey := fmt.Sprintf("makan:%d", id)
	var cached dto.MakanResponse
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

func (s *makanService) Update(id uint, req dto.UpdateMakanRequest) (*dto.MakanResponse, error) {
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

	s.cache.Delete(fmt.Sprintf("makan:%d", id))
	s.cache.DeletePattern("makan:*")

	
	logger.AuditLogger.Info().
		Uint("makan_id", entity.ID).
		Str("action", "makan_update").
		Msg("makan updated")
	

	return s.mapToResponse(entity), nil
}

func (s *makanService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("makan:%d", id))
	s.cache.DeletePattern("makan:*")

	
	logger.AuditLogger.Info().
		Uint("makan_id", id).
		Str("action", "makan_deletion").
		Msg("makan deleted")
	

	return s.repo.Delete(id)
}

func (s *makanService) mapToResponse(entity *entity.Makan) *dto.MakanResponse {
	return &dto.MakanResponse{
		ID:        entity.ID,
		Name: entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
