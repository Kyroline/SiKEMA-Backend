package excuseHandler

import (
	updateExcuse "attendance-is/controllers/excuse/student/update"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateExcuse.Service
}

func NewUpdateExcuseHandler(service updateExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateExcuseHandler(c *gin.Context) {
	var input updateExcuse.InputUpdateExcuse
	c.ShouldBindJSON(&input)

	_, err := h.service.UpdateExcuseService(&input)

	switch err {
	case "UPDATE_EXCUSE_NOTFOUND_404":
		return
	case "UPDATE_EXCUSE_INTERNAL_500":
		return
	default:
		c.Status(http.StatusAccepted)
		return
	}
}
