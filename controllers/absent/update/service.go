package updateAbsent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateAbsent(c *gin.Context) {
	var input Input
	c.ShouldBindJSON(&input)

	id := c.Param("id")
	var absent model.Absent
	if err := util.DB.Where("id = ?", id).Find(&absent).Error; util.HandleGORMError(c, err, http.StatusNotFound, "") {
		return
	}

	absent.Hours = input.Hours

	if err := util.DB.Save(&absent).Error; util.HandleGORMError(c, err, http.StatusInternalServerError, "") {
		return
	}

	c.Status(http.StatusNoContent)
}
