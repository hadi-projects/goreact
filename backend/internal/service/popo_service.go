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

type PopoService interface {
	Create(req dto.CreatePopoRequest) (*dto.PopoResponse, error)
	GetAll(pagination *defaultdto.PaginationRequest) (*defaultdto.PaginationResponse, error)
	GetByID(id uint) (*dto.PopoResponse, error)
	Update(id uint, req dto.UpdatePopoRequest) (*dto.PopoResponse, error)
	Delete(id uint) error
}

type popoService struct {
	repo  repository.PopoRepository
	cache cache.CacheService
}

func NewPopoService(repo repository.PopoRepository, cache cache.CacheService) PopoService {
	return &popoService{
		repo:  repo,
		cache: cache,
	}
}

func (s *popoService) Create(req dto.CreatePopoRequest) (*dto.PopoResponse, error) {
	entity := &entity.Popo{
		Name: req.Name,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("popo:*")

	
	logger.AuditLogger.Info().
		Uint("popo_id", entity.ID).
		Str("action", "popo_creation").
		Msg("popo created")
	

	return s.mapToResponse(entity), nil
}

func (s *popoService) GetAll(pagination *defaultdto.PaginationRequest) (*defaultdto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("popo:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached defaultdto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.PopoResponse
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

func (s *popoService) GetByID(id uint) (*dto.PopoResponse, error) {
	cacheKey := fmt.Sprintf("popo:%d", id)
	var cached dto.PopoResponse
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

func (s *popoService) Update(id uint, req dto.UpdatePopoRequest) (*dto.PopoResponse, error) {
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

	s.cache.Delete(fmt.Sprintf("popo:%d", id))
	s.cache.DeletePattern("popo:*")

	
	logger.AuditLogger.Info().
		Uint("popo_id", entity.ID).
		Str("action", "popo_update").
		Msg("popo updated")
	

	return s.mapToResponse(entity), nil
}

func (s *popoService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("popo:%d", id))
	s.cache.DeletePattern("popo:*")

	
	logger.AuditLogger.Info().
		Uint("popo_id", id).
		Str("action", "popo_deletion").
		Msg("popo deleted")
	

	return s.repo.Delete(id)
}

func (s *popoService) mapToResponse(entity *entity.Popo) *dto.PopoResponse {
	return &dto.PopoResponse{
		ID:        entity.ID,
		Name: entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
