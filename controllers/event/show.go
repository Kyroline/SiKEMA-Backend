package event

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) FindEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	event, _ := h.service.FindEvent(uint(id))

	c.JSON(http.StatusOK, gin.H{
		"data": event,
	})
}
