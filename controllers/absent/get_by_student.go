package absent

import (
	schema "attendance-is/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAbsentsByStudent(c *gin.Context) {
	meta := schema.Metadata{Page: 1, ItemPerPage: 10}

	if c.Query("page") != "" {
		meta.Page, _ = strconv.Atoi(c.Query("page"))
	}

	if c.Query("itemperpage") != "" {
		meta.ItemPerPage, _ = strconv.Atoi(c.Query("itemperpage"))
	}

	studentID, _ := strconv.Atoi(c.Param("studentid"))

	absents, err := h.service.GetAbsentsByStudent(&meta, studentID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": absents,
		"meta": meta,
	})

}

func (h *handler) GetAbsentByStudent(c *gin.Context) {
	meta := schema.Metadata{Page: 1, ItemPerPage: 10}

	if c.Query("page") != "" {
		meta.Page, _ = strconv.Atoi(c.Query("page"))
	}

	if c.Query("itemperpage") != "" {
		meta.ItemPerPage, _ = strconv.Atoi(c.Query("itemperpage"))
	}

	studentID, _ := strconv.Atoi(c.Param("studentid"))

	ID, _ := strconv.Atoi(c.Param("id"))

	absents, err := h.service.GetAbsentByStudent(&meta, studentID, ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": absents,
		"meta": meta,
	})

}
