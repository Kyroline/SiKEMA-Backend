package getEventHandler

import (
	getEvent "attendance-is/controllers/event/lecturer/get"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service getEvent.Service
}

func NewGetEventHandler(service getEvent.Service) *handler {
	return &handler{service: service}
}

func (h *handler) GetEventHandler(c *gin.Context) {
	res, err := h.service.GetEventService()
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
