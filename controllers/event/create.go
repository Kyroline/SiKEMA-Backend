package event

import (
	schema "attendance-is/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateEventHandler(c *gin.Context) {
	var input schema.CreateEventRequest
	c.ShouldBindJSON(&input)

	newEvent, err := h.service.CreateEvent(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newEvent,
	})
}
