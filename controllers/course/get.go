package course

import (
	schema "attendance-is/schemas"
	service "attendance-is/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetCourseByLecturer(c *gin.Context) {
	input := service.InputGetCourseByLecturer{LecturerID: c.Param("lecturerid")}
	courses, err := h.service.GetCourseByLecturer(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courses,
	})
	return
}

func (h *handler) GetCourseByPBM(c *gin.Context) {
	meta := schema.Metadata{Page: 1, ItemPerPage: 10}

	if c.Query("page") != "" {
		meta.Page, _ = strconv.Atoi(c.Query("page"))
	}

	if c.Query("itemperpage") != "" {
		meta.ItemPerPage, _ = strconv.Atoi(c.Query("itemperpage"))
	}

	courses, err := h.service.GetCourseByPBM(&meta)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": courses,
	})
	return
}
