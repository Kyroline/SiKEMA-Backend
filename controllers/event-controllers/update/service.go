package updateEvent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")

	var input Input
	c.ShouldBindJSON(&input)

	var event model.Event
	if err := util.DB.Where("id = ?", id).Find(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	event.Status = input.Status
	if err := util.DB.Save(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusAccepted)
}
