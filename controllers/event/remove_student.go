package event

import (
	schema "attendance-is/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) RemoveStudent(c *gin.Context) {
	var data schema.AddStudentToEventRequest
	c.ShouldBindJSON(&data)
	id := c.Param("id")
	eventId, _ := strconv.ParseUint(id, 0, 64)
	data.EventId = uint(eventId)

	_, err := h.service.RemoveStudent(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.Status(http.StatusNoContent)
}
