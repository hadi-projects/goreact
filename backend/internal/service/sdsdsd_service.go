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

type SdsdsdService interface {
	Create(req dto.CreateSdsdsdRequest) (*dto.SdsdsdResponse, error)
	GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error)
	GetByID(id uint) (*dto.SdsdsdResponse, error)
	Update(id uint, req dto.UpdateSdsdsdRequest) (*dto.SdsdsdResponse, error)
	Delete(id uint) error
}

type sdsdsdService struct {
	repo  repository.SdsdsdRepository
	cache cache.CacheService
}

func NewSdsdsdService(repo repository.SdsdsdRepository, cache cache.CacheService) SdsdsdService {
	return &sdsdsdService{
		repo:  repo,
		cache: cache,
	}
}

func (s *sdsdsdService) Create(req dto.CreateSdsdsdRequest) (*dto.SdsdsdResponse, error) {
	entity := &entity.Sdsdsd{
		Name: req.Name,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("sdsdsdsdd:*")

	
	logger.AuditLogger.Info().
		Uint("sdsdsd_id", entity.ID).
		Str("action", "sdsdsd_creation").
		Msg("sdsdsd created")
	

	return s.mapToResponse(entity), nil
}

func (s *sdsdsdService) GetAll(pagination *dto.PaginationRequest) (*dto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("sdsdsdsdd:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached dto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.SdsdsdResponse
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

func (s *sdsdsdService) GetByID(id uint) (*dto.SdsdsdResponse, error) {
	cacheKey := fmt.Sprintf("sdsdsdsdd:%d", id)
	var cached dto.SdsdsdResponse
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

func (s *sdsdsdService) Update(id uint, req dto.UpdateSdsdsdRequest) (*dto.SdsdsdResponse, error) {
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

	s.cache.Delete(fmt.Sprintf("sdsdsdsdd:%d", id))
	s.cache.DeletePattern("sdsdsdsdd:*")

	
	logger.AuditLogger.Info().
		Uint("sdsdsd_id", entity.ID).
		Str("action", "sdsdsd_update").
		Msg("sdsdsd updated")
	

	return s.mapToResponse(entity), nil
}

func (s *sdsdsdService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("sdsdsdsdd:%d", id))
	s.cache.DeletePattern("sdsdsdsdd:*")

	
	logger.AuditLogger.Info().
		Uint("sdsdsd_id", id).
		Str("action", "sdsdsd_deletion").
		Msg("sdsdsd deleted")
	

	return s.repo.Delete(id)
}

func (s *sdsdsdService) mapToResponse(entity *entity.Sdsdsd) *dto.SdsdsdResponse {
	return &dto.SdsdsdResponse{
		ID:        entity.ID,
		Name: entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
