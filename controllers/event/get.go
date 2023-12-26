package event

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetEvent(c *gin.Context) {
	var result interface{}
	if c.Query("course_id") == "" && c.Query("class_id") == "" {
		result, _ = h.service.GetEvent()
	} else if c.Query("course_id") != "" && c.Query("class_id") == "" {
		course_id := c.Query("course_id")
		result, _ = h.service.GetEventByWhere("course_id = ?", course_id)
	} else if c.Query("course_id") == "" && c.Query("class_id") != "" {
		class_id := c.Query("class_id")
		result, _ = h.service.GetEventByWhere("class_id = ?", class_id)
	} else {
		course_id := c.Query("course_id")
		class_id := c.Query("class_id")
		result, _ = h.service.GetEventByWhere("course_id = ? AND class_id = ?", course_id, class_id)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
	return
}
