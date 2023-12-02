package getExcuseHandler

import (
	getExcuse "attendance-is/controllers/excuse/student/get"
	"net/http"

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
	res, err := h.service.GetExcuseService(input)

	switch err {
	case "ERR":

	default:
		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	}
}

// func (h *handler) GetExcuseHandler(c *gin.Context) {
// 	var input getExcuse.InputGetExcuseByStudent
// 	input.StudentID = c.Param("studentid")
// 	input.ID = c.Param("id")
// 	res, err := h.service.GetExcuseByStudentService(input)

// 	switch err {
// 	case "ERR":

// 	default:
// 		c.JSON(http.StatusOK, gin.H{
// 			"data": res,
// 		})
// 	}
// }
