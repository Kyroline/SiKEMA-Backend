package createEvent

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	model "attendance-is/models"
	util "attendance-is/utils"
)

func CreateEvent(c *gin.Context) {
	var input Input
	c.ShouldBindJSON(&input)

	newEvent := model.Event{
		Model:   gorm.Model{ID: input.Id},
		ClassId: input.ClassId,
		Meet:    input.Meet,
		Status:  0, // 0 Means open
	}

	// newEvent.LecturerId = get lecturer id from token

	if err := util.DB.Model(model.Event{}).Create(&newEvent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newEvent,
	})
}
