package course

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetCourseByLecturer(c *gin.Context) {
	courses, err := h.service.GetCourseByLecturer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courses,
	})
	return
}

func (h *handler) GetCourseByPBM(c *gin.Context) {
	courses, err := h.service.GetCourseByLecturer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courses,
	})
	return
}
