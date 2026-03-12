package service

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math"
	"time"

	"github.com/xuri/excelize/v2"
	defaultDto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"github.com/hadi-projects/go-react-starter/internal/repository"
	"github.com/hadi-projects/go-react-starter/pkg/cache"
	"github.com/hadi-projects/go-react-starter/pkg/logger"
)

type MinaService interface {
	Create(req dto.CreateMinaRequest) (*dto.MinaResponse, error)
	GetAll(pagination *defaultDto.PaginationRequest) (*defaultDto.PaginationResponse, error)
	GetByID(id uint) (*dto.MinaResponse, error)
	Update(id uint, req dto.UpdateMinaRequest) (*dto.MinaResponse, error)
	Delete(id uint) error
	Export(format string) ([]byte, error)
}

type minaService struct {
	repo  repository.MinaRepository
	cache cache.CacheService
}

func NewMinaService(repo repository.MinaRepository, cache cache.CacheService) MinaService {
	return &minaService{
		repo:  repo,
		cache: cache,
	}
}

func (s *minaService) Create(req dto.CreateMinaRequest) (*dto.MinaResponse, error) {
	entity := &entity.Mina{
		Name: req.Name,
	}

	if err := s.repo.Create(entity); err != nil {
		return nil, err
	}

	s.cache.DeletePattern("mina:*")

	
	logger.AuditLogger.Info().
		Uint("mina_id", entity.ID).
		Str("action", "mina_creation").
		Msg("mina created")
	

	return s.mapToResponse(entity), nil
}

func (s *minaService) GetAll(pagination *defaultDto.PaginationRequest) (*defaultDto.PaginationResponse, error) {
	cacheKey := fmt.Sprintf("mina:page:%d:limit:%d:search:%s", pagination.GetPage(), pagination.GetLimit(), pagination.Search)
	var cached defaultDto.PaginationResponse
	if err := s.cache.Get(cacheKey, &cached); err == nil {
		return &cached, nil
	}

	entities, total, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	var responses []dto.MinaResponse
	for _, e := range entities {
		responses = append(responses, *s.mapToResponse(&e))
	}

	response := &defaultDto.PaginationResponse{
		Data: responses,
		Meta: defaultDto.PaginationMeta{
			CurrentPage: pagination.GetPage(),
			TotalPages:  int(math.Ceil(float64(total) / float64(pagination.GetLimit()))),
			TotalData:   total,
			Limit:       pagination.GetLimit(),
		},
	}

	s.cache.Set(cacheKey, response, 5*time.Minute)
	return response, nil
}

func (s *minaService) GetByID(id uint) (*dto.MinaResponse, error) {
	cacheKey := fmt.Sprintf("mina:%d", id)
	var cached dto.MinaResponse
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

func (s *minaService) Update(id uint, req dto.UpdateMinaRequest) (*dto.MinaResponse, error) {
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

	s.cache.Delete(fmt.Sprintf("mina:%d", id))
	s.cache.DeletePattern("mina:*")

	
	logger.AuditLogger.Info().
		Uint("mina_id", entity.ID).
		Str("action", "mina_update").
		Msg("mina updated")
	

	return s.mapToResponse(entity), nil
}

func (s *minaService) Delete(id uint) error {
	s.cache.Delete(fmt.Sprintf("mina:%d", id))
	s.cache.DeletePattern("mina:*")

	
	logger.AuditLogger.Info().
		Uint("mina_id", id).
		Str("action", "mina_deletion").
		Msg("mina deleted")
	

	return s.repo.Delete(id)
}

func (s *minaService) Export(format string) ([]byte, error) {
	pagination := &defaultDto.PaginationRequest{
		Page:  1,
		Limit: 1000000,
	}

	entities, _, err := s.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	if format == "csv" {
		return s.generateCSV(entities)
	}
	return s.generateExcel(entities)
}

func (s *minaService) generateCSV(entities []entity.Mina) ([]byte, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	header := []string{"ID", "Name", "Created At"}
	writer.Write(header)

	for _, e := range entities {
		row := []string{
			fmt.Sprintf("%d", e.ID),
			fmt.Sprintf("%v", e.Name),
			e.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		writer.Write(row)
	}

	writer.Flush()
	return buf.Bytes(), nil
}

func (s *minaService) generateExcel(entities []entity.Mina) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "Sheet1"
	f.SetSheetName("Sheet1", sheet)

	header := []string{"ID", "Name", "Created At"}
	for i, h := range header {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	for i, e := range entities {
		rowNum := i + 2
		f.SetCellValue(sheet, fmt.Sprintf("A%d", rowNum), e.ID)
		cell, _ := excelize.CoordinatesToCellName(0+2, rowNum)
		f.SetCellValue(sheet, cell, e.Name)
		lastCell, _ := excelize.CoordinatesToCellName(len(header), rowNum)
		f.SetCellValue(sheet, lastCell, e.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *minaService) mapToResponse(entity *entity.Mina) *dto.MinaResponse {
	return &dto.MinaResponse{
		ID:        entity.ID,
		Name: entity.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
