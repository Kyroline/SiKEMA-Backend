package class

import (
	schema "attendance-is/schemas/class"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateClassHandler(c *gin.Context) {
	data := schema.UpdateClassRequest{}
	c.ShouldBindJSON(&data)

	id, _ := strconv.ParseUint(c.Param("id"), 0, 16)
	data.ID = uint(id)

	_, err := h.service.UpdateClass(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
