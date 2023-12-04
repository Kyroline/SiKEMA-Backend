package getExcuseHandler

import (
	getExcuse "attendance-is/controllers/excuse/pbm/get"
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
	ID, _ := strconv.Atoi(c.Param("id"))
	input.ID = uint(ID)

	res, err := h.service.GetExcuseService(input)
	switch err {
	default:
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}
