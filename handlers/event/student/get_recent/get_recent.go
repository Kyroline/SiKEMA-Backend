package getRecentEventHandler

import (
	getRecentEvent "attendance-is/controllers/event/student/get_recent"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getRecentEvent.Service
}

func NewGetRecentEventHandler(service getRecentEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetRecentEventHandler(c *gin.Context) {
	input := getRecentEvent.InputGetRecentEvent{StudentID: c.Param("studentid")}
	res, err := h.service.GetRecentEventService(input)
	if err != "" {
		switch err {
		case "EVENT_NOTFOUND_404":
			util.ErrorRespose(c, http.StatusNotFound, "Record not found")
			return
		default:
			util.ErrorRespose(c, http.StatusInternalServerError, err)
			return
		}
	}
	util.APIResponse(c, http.StatusOK, res, nil)
}
