package absent

import (
	schema "attendance-is/schemas"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAbsentsByPBM(c *gin.Context) {
	meta := schema.Metadata{Page: 1, ItemPerPage: 10}

	if c.Query("page") != "" {
		meta.Page, _ = strconv.Atoi(c.Query("page"))
	}

	if c.Query("itemperpage") != "" {
		meta.ItemPerPage, _ = strconv.Atoi(c.Query("itemperpage"))
	}

	absent, err := h.service.GetAbsentsByPBM(&meta)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"metadata": meta,
		"data":     absent,
	})
}

func (h *handler) GetAbsentByPBM(c *gin.Context) {
	meta := schema.Metadata{Page: 1, ItemPerPage: 10}

	if c.Query("page") != "" {
		meta.Page, _ = strconv.Atoi(c.Query("page"))
	}

	if c.Query("itemperpage") != "" {
		meta.ItemPerPage, _ = strconv.Atoi(c.Query("itemperpage"))
	}

	id, _ := strconv.Atoi(c.Param("id"))

	absent, err := h.service.GetAbsentByPBM(&meta, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"metadata": meta,
		"data":     absent,
	})
}
