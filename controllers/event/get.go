package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetEvent(c *gin.Context) {
	event, _ := h.service.GetEvent()

	c.JSON(http.StatusOK, gin.H{
		"data": event,
	})
}
