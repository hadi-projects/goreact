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

type AkusajaService interface {
	Create(req dto.CreateAkusajaRequest) (*dto.AkusajaResponse, error)
	GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	GetByID(id uint) (*dto.AkusajaResponse, error)
	Update(id uint, req dto.UpdateAkusajaRequest) (*dto.AkusajaResponse, error)
	Delete(id uint) error
}

type akusajaService struct {
	repo  repository.AkusajaRepository
	cache cache.CacheService
}

func NewAkusajaService(repo repository.AkusajaRepository, cache cache.CacheService) AkusajaService {
	return &akusajaService{
		repo:  repo,
		cache: cache,
	}
}

func (s *akusajaService) Create(req dto.CreateAkusajaRequest) (*dto.AkusajaResponse, error) {
	entity := &entity.Akusaja{
		Name: req.Name,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("akusaja:*")

	
	logger.AuditLogger.Info().
		Uint("akusaja_id", entity.ID).
		Str("action", "akusaja_creation").
		Msg("akusaja created")
	

	return s.mapToResponse(entity), nil
}

func (s *akusajaService) GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("akusaja:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached dto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.AkusajaResponse
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

func (s *akusajaService) GetByID(id uint) (*dto.AkusajaResponse, error) {
	cacheKey := fmt.Sprintf("akusaja:%d", id)
	var cached dto.AkusajaResponse
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

func (s *akusajaService) Update(id uint, req dto.UpdateAkusajaRequest) (*dto.AkusajaResponse, error) {
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

	s.cache.Delete(fmt.Sprintf("akusaja:%d", id))
	s.cache.DeletePattern("akusaja:*")

	
	logger.AuditLogger.Info().
		Uint("akusaja_id", entity.ID).
		Str("action", "akusaja_update").
		Msg("akusaja updated")
	

	return s.mapToResponse(entity), nil
}

func (s *akusajaService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("akusaja:%d", id))
	s.cache.DeletePattern("akusaja:*")

	
	logger.AuditLogger.Info().
		Uint("akusaja_id", id).
		Str("action", "akusaja_deletion").
		Msg("akusaja deleted")
	

	return s.repo.Delete(id)
}

func (s *akusajaService) mapToResponse(entity *entity.Akusaja) *dto.AkusajaResponse {
	return &dto.AkusajaResponse{
		ID:        entity.ID,
		Name: entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
