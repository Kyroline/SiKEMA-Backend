package createAbsent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAbsent(c *gin.Context) {
	var input Input
	c.ShouldBindJSON(&input)

	newAbsent := model.Absent{
		StudentID: input.StudentId,
		Hours:     input.Hours,
	}

	if err := util.DB.Create(&newAbsent).Error; util.HandleGORMError(c, err, http.StatusInternalServerError, "") {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newAbsent,
	})
}

func CreateAbsentWithEvent(c *gin.Context, event model.Event) {
	var input Input
	c.ShouldBindJSON(&input)

	newAbsent := model.Absent{
		StudentID: input.StudentId,
		Hours:     input.Hours,
		Event:     event,
	}

	if err := util.DB.Create(&newAbsent).Error; util.HandleGORMError(c, err, http.StatusInternalServerError, "") {
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newAbsent,
	})
}
