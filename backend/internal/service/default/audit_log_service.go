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

type AuditLogService interface {
	Create(log *entity.AuditLog) error
	GetAll(query *dto.AuditLogQuery) ([]dto.AuditLogResponse, int64, error)
}

type auditLogService struct {
	repo  repository.AuditLogRepository
	cache cache.CacheService
}

func NewAuditLogService(repo repository.AuditLogRepository, cache cache.CacheService) AuditLogService {
	return &auditLogService{
		repo:  repo,
		cache: cache,
	}
}

func (s *auditLogService) Create(log *entity.AuditLog) error {
	return s.repo.Create(&logger.AuditLog{
		RequestID: log.RequestID,
		UserID:    log.UserID,
		UserEmail: log.UserEmail,
		Action:    log.Action,
		Module:    log.Module,
		TargetID:  log.TargetID,
		Metadata:  log.Metadata,
	})
}

func (s *auditLogService) GetAll(query *dto.AuditLogQuery) ([]dto.AuditLogResponse, int64, error) {
	cacheKey := fmt.Sprintf("audit_logs:%d:%d:%s:%s:%s:%s",
		query.GetPage(),
		query.GetLimit(),
		query.Module,
		query.Action,
		query.UserEmail,
		query.RequestID,
	)

	type cacheData struct {
		Responses []dto.AuditLogResponse `json:"responses"`
		Total     int64                  `json:"total"`
	}

	var cached cacheData
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return cached.Responses, cached.Total, nil
	}

	logs, total, err := s.repo.FindAll(query)
	if err != nil {
		return nil, 0, err
	}

	var responses []dto.AuditLogResponse
	for _, l := range logs {
		responses = append(responses, dto.AuditLogResponse{
			ID:        l.ID,
			RequestID: l.RequestID,
			UserID:    l.UserID,
			UserEmail: l.UserEmail,
			Action:    l.Action,
			Module:    l.Module,
			TargetID:  l.TargetID,
			Metadata:  l.Metadata,
			CreatedAt: l.CreatedAt,
		})
	}

	_ = s.cache.Set(cacheKey, cacheData{
		Responses: responses,
		Total:     total,
	}, 10*time.Second)

	return responses, total, nil
}
