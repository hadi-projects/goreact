package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	service "github.com/hadi-projects/go-react-starter/internal/service/default"
)

type AuditLogHandler interface {
	GetAll(c *gin.Context)
}

type auditLogHandler struct {
	service service.AuditLogService
}

func NewAuditLogHandler(service service.AuditLogService) AuditLogHandler {
	return &auditLogHandler{service: service}
}

func (h *auditLogHandler) GetAll(c *gin.Context) {
	var query dto.AuditLogQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logs, total, err := h.service.GetAll(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  logs,
		"total": total,
		"page":  query.GetPage(),
		"limit": query.GetLimit(),
	})
}
