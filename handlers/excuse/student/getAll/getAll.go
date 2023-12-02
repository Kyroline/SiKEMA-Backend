package getExcuseHandler

import (
	getAllExcuse "attendance-is/controllers/excuse/student/getAll"
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
	var input getAllExcuse.InputGetAllExcuse
	input.StudentID = c.Param("studentid")
	res, err := h.service.GetAllExcuseService(input)

	switch err {
	case "ERR":

	default:
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}

// }
