package getAbsentHandler

import (
	getAbsent "attendance-is/controllers/absent/student/get"
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getAbsent.Service
}

func NewGetAbsentHandler(service getAbsent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetAbsentHandler(c *gin.Context) {
	var input getAbsent.InputGetAbsent
	input.StudentID = c.Param("studentid")

	res, err := h.service.GetAbsentService(input)
	if err != "" {
		switch err {
		case "ABSENT_NOTFOUND_404":
			util.ErrorRespose(c, http.StatusNotFound, "Record not found")
			return
		default:
			util.ErrorRespose(c, http.StatusInternalServerError, err)
			return
		}
	}
	var result []model.AbsentJSON
	for _, a := range *res {
		result = append(result, *a.ParseJSON())
	}
	util.APIResponse(c, http.StatusOK, result, nil)
}
