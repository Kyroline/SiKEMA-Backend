package getExcuseHandler

import (
	getExcuse "attendance-is/controllers/excuse/student/get"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getExcuse.Service
}

func NewGetExcuseHandler(service getExcuse.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetExcuseHandler(c *gin.Context) {
	var input getExcuse.InputGetExcuse
	input.StudentID = c.Param("studentid")

	id, _ := strconv.Atoi(c.Param("id"))
	input.ID = uint(id)
	res, err := h.service.GetExcuseService(input)

	switch err {
	case "ERR":

	default:
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}

func (h *handler) GetExcusesHandler(c *gin.Context) {
	var input getExcuse.InputGetExcuse
	input.StudentID = c.Param("studentid")

	id, _ := strconv.Atoi(c.Param("id"))
	input.ID = uint(id)
	res, err := h.service.GetExcusesService(input)

	switch err {
	case "ERR":

	default:
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
