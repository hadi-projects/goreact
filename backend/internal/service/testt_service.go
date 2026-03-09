package service

import (
	"fmt"
	"math"
	"time"

	"github.com/hadi-projects/go-react-starter/internal/dto"
	defaultdto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"github.com/hadi-projects/go-react-starter/internal/repository"
	"github.com/hadi-projects/go-react-starter/pkg/cache"
	"github.com/hadi-projects/go-react-starter/pkg/logger"
)

type TesttService interface {
	Create(req dto.CreateTesttRequest) (*dto.TesttResponse, error)
	GetAll(pagination *defaultdto.PaginationRequest) (*defaultdto.PaginationResponse, error)
	GetByID(id uint) (*dto.TesttResponse, error)
	Update(id uint, req dto.UpdateTesttRequest) (*dto.TesttResponse, error)
	Delete(id uint) error
}

type testtService struct {
	repo  repository.TesttRepository
	cache cache.CacheService
}

func NewTesttService(repo repository.TesttRepository, cache cache.CacheService) TesttService {
	return &testtService{
		repo:  repo,
		cache: cache,
	}
}

func (s *testtService) Create(req dto.CreateTesttRequest) (*dto.TesttResponse, error) {
	entity := &entity.Testt{
		Name: req.Name,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("testt:*")

	
	logger.AuditLogger.Info().
		Uint("testt_id", entity.ID).
		Str("action", "testt_creation").
		Msg("testt created")
	

	return s.mapToResponse(entity), nil
}

func (s *testtService) GetAll(pagination *defaultdto.PaginationRequest) (*defaultdto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("testt:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached defaultdto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.TesttResponse
	for _, e := range entities {
		responses = append(responses, *s.mapToResponse(&e))
	}

	response := &defaultdto.PaginationResponse{
		Data: responses,
		Meta: defaultdto.PaginationMeta{
			CurrentPage: pagination.GetPage(),
			TotalPages:  int(math.Ceil(float64(total) / float64(pagination.GetLimit()))),
			TotalData:   total,
			Limit:       pagination.GetLimit(),
		},
	}

	s.cache.Set(cacheKey, response, 5*time.Minute)
	return response, nil
}

func (s *testtService) GetByID(id uint) (*dto.TesttResponse, error) {
	cacheKey := fmt.Sprintf("testt:%d", id)
	var cached dto.TesttResponse
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

func (s *testtService) Update(id uint, req dto.UpdateTesttRequest) (*dto.TesttResponse, error) {
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

	s.cache.Delete(fmt.Sprintf("testt:%d", id))
	s.cache.DeletePattern("testt:*")

	
	logger.AuditLogger.Info().
		Uint("testt_id", entity.ID).
		Str("action", "testt_update").
		Msg("testt updated")
	

	return s.mapToResponse(entity), nil
}

func (s *testtService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("testt:%d", id))
	s.cache.DeletePattern("testt:*")

	
	logger.AuditLogger.Info().
		Uint("testt_id", id).
		Str("action", "testt_deletion").
		Msg("testt deleted")
	

	return s.repo.Delete(id)
}

func (s *testtService) mapToResponse(entity *entity.Testt) *dto.TesttResponse {
	return &dto.TesttResponse{
		ID:        entity.ID,
		Name: entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
