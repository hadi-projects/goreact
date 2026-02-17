package service

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hadi-projects/go-react-starter/config"
	"github.com/hadi-projects/go-react-starter/internal/dto"
)

type LogService interface {
	GetLogs(query dto.LogQuery) ([]dto.LogResponse, error)
}

type logService struct {
	config *config.Config
}

func NewLogService(config *config.Config) LogService {
	return &logService{config: config}
}

func (s *logService) GetLogs(query dto.LogQuery) ([]dto.LogResponse, error) {
	var allLogs []dto.LogResponse

	filesToRead := []string{}
	if query.Type == "" {
		query.Type = "all"
	}
	if query.Type == "all" || query.Type == "auth" {
		filesToRead = append(filesToRead, "auth.log")
	}
	if query.Type == "all" || query.Type == "audit" {
		filesToRead = append(filesToRead, "audit.log")
	}
	if query.Type == "all" || query.Type == "system" {
		filesToRead = append(filesToRead, "system.log")
	}

	for _, fileName := range filesToRead {
		filePath := filepath.Join(s.config.Log.Dir, fileName)
		logs, err := s.readLogFile(filePath, strings.TrimSuffix(fileName, ".log"))
		if err != nil {
			// Skip if file doesn't exist yet
			continue
		}
		allLogs = append(allLogs, logs...)
	}

	// Filter by UserID if provided
	if query.UserID != 0 {
		var filteredLogs []dto.LogResponse
		for _, log := range allLogs {
			if log.UserID != nil && *log.UserID == query.UserID {
				filteredLogs = append(filteredLogs, log)
			}
		}
		allLogs = filteredLogs
	}

	// Sort logs by time descending
	sort.Slice(allLogs, func(i, j int) bool {
		return allLogs[i].Time.After(allLogs[j].Time)
	})

	return allLogs, nil
}

func (s *logService) readLogFile(filePath string, logType string) ([]dto.LogResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []dto.LogResponse
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		var raw map[string]interface{}
		if err := json.Unmarshal([]byte(line), &raw); err != nil {
			continue
		}

		log := dto.LogResponse{
			Type: logType,
		}

		if val, ok := raw["level"].(string); ok {
			log.Level = val
		}
		if val, ok := raw["action"].(string); ok {
			log.Action = val
		}
		if val, ok := raw["message"].(string); ok {
			log.Message = val
		}
		if val, ok := raw["email"].(string); ok {
			log.Email = val
		}
		if val, ok := raw["request_id"].(string); ok {
			log.RequestID = val
		}

		// Handle user_id (it could be uint or float64 from json)
		if val, ok := raw["user_id"]; ok {
			switch v := val.(type) {
			case float64:
				u := uint(v)
				log.UserID = &u
			case uint:
				log.UserID = &v
			}
		}

		if val, ok := raw["time"].(string); ok {
			if t, err := json.Marshal(val); err == nil {
				json.Unmarshal(t, &log.Time)
			}
			// Zerolog default time format is often RFC3339 or similar
			// dto.LogResponse has time.Time, so Unmarshal handles it if it's JSON string
		}

		// Collect other fields into Details
		delete(raw, "level")
		delete(raw, "action")
		delete(raw, "message")
		delete(raw, "user_id")
		delete(raw, "email")
		delete(raw, "time")
		delete(raw, "request_id")
		log.Details = raw

		logs = append(logs, log)
	}

	return logs, nil
}
