package event

import (
	schema "attendance-is/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) FinalizeEventHandler(c *gin.Context) {
	id := c.Param("id")
	eventId, _ := strconv.ParseUint(id, 0, 64)

	event, err := h.service.FinalizeEvent(schema.FinalizeEventRequest{EventId: uint(eventId)})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": event,
	})
}
