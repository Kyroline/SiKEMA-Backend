package createExcuseHandler

import (
	createExcuse "attendance-is/controllers/excuse/student/create"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service createExcuse.Service
}

func NewCreateExcuseHandler(service createExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateExcuseHandler(c *gin.Context) {
	var input createExcuse.InputCreateExcuse
	c.ShouldBindJSON(&input)

	res, err := h.service.CreateExcuseService(&input)

	switch err {
	case "CREATE_EXCUSE_CONFLICT_409":
		c.JSON(http.StatusConflict, gin.H{
			"message": "Excuse ID already exist",
		})
		return
	case "CREATE_EXCUSE_INTERNAL_500":
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "There's an error creating a new resource",
		})
		return
	default:
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
		return
	}
}
