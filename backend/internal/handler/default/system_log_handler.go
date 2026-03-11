package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/hadi-projects/go-react-starter/internal/dto/default"
	service "github.com/hadi-projects/go-react-starter/internal/service/default"
	"github.com/hadi-projects/go-react-starter/pkg/response"
)

type SystemLogHandler interface {
	GetAll(ctx *gin.Context)
}

type systemLogHandler struct {
	service service.SystemLogService
}

func NewSystemLogHandler(service service.SystemLogService) SystemLogHandler {
	return &systemLogHandler{service: service}
}

func (h *systemLogHandler) GetAll(ctx *gin.Context) {
	var query dto.SystemLogQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid query parameters")
		return
	}

	logs, total, err := h.service.GetAll(&query)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "Failed to get system logs")
		return
	}

	limit := query.GetLimit()
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	response.SuccessWithPagination(ctx, http.StatusOK, "System logs retrieved successfully", logs, &response.PaginationMeta{
		CurrentPage: query.GetPage(),
		TotalPages:  totalPages,
		TotalData:   total,
		Limit:       limit,
	})
}
