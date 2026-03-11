package service

import (
	"fmt"
	"time"

	dto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	entity "github.com/hadi-projects/go-react-starter/internal/entity/default"
	repository "github.com/hadi-projects/go-react-starter/internal/repository/default"
	"github.com/hadi-projects/go-react-starter/pkg/cache"
	"github.com/hadi-projects/go-react-starter/pkg/logger"
)

type SystemLogService interface {
	Create(log *entity.SystemLog) error
	GetAll(query *dto.SystemLogQuery) ([]dto.SystemLogResponse, int64, error)
}

type systemLogService struct {
	repo  repository.SystemLogRepository
	cache cache.CacheService
}

func NewSystemLogService(repo repository.SystemLogRepository, cache cache.CacheService) SystemLogService {
	return &systemLogService{
		repo:  repo,
		cache: cache,
	}
}

func (s *systemLogService) Create(log *entity.SystemLog) error {
	return s.repo.Create(&logger.SystemLog{
		RequestID:    log.RequestID,
		Method:       log.Method,
		Path:         log.Path,
		StatusCode:   log.StatusCode,
		Latency:      log.Latency,
		RequestBody:  log.RequestBody,
		ResponseBody: log.ResponseBody,
	})
}

func (s *systemLogService) GetAll(query *dto.SystemLogQuery) ([]dto.SystemLogResponse, int64, error) {
	// Try to get from cache
	cacheKey := fmt.Sprintf("system_logs:%d:%d:%s:%s:%d:%s",
		query.GetPage(),
		query.GetLimit(),
		query.Method,
		query.Path,
		query.StatusCode,
		query.RequestID,
	)

	type cacheData struct {
		Responses []dto.SystemLogResponse `json:"responses"`
		Total     int64                   `json:"total"`
	}

	var cached cacheData
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return cached.Responses, cached.Total, nil
	}

	logs, total, err := s.repo.FindAll(query)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.SystemLogResponse
	for _, l := range logs {
		responses = append(responses, dto.SystemLogResponse{
			ID:           l.ID,
			RequestID:    l.RequestID,
			Method:       l.Method,
			Path:         l.Path,
			StatusCode:   l.StatusCode,
			Latency:      l.Latency,
			RequestBody:  l.RequestBody,
			ResponseBody: l.ResponseBody,
			CreatedAt:    l.CreatedAt,
		})
	}

	// Save to cache with short TTL (10 seconds)
	_ = s.cache.Set(cacheKey, cacheData{
		Responses: responses,
		Total:     total,
	}, 10*time.Second)

	return responses, total, nil
}
