package showAbsent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAbsent(c *gin.Context) {
	var absent []model.Absent
	if err := util.DB.Find(&absent).Error; util.HandleGORMError(c, err, http.StatusNotFound, "") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": absent,
	})
}

func ShowAbsent(c *gin.Context) {
	id := c.Param("id")
	var absent model.Absent
	if err := util.DB.Where("id = ?", id).Find(&absent).Error; util.HandleGORMError(c, err, http.StatusNotFound, "") {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": absent,
	})
}
