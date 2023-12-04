package getAllExcuseHandler

import (
	getAllExcuse "attendance-is/controllers/excuse/pbm/getAll"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAllExcuse.Service
}

func NewGetAllExcuseHandler(service getAllExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAllExcuseHandler(c *gin.Context) {
	res, err := h.service.GetAllExcuseService()
	switch err {
	default:
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
