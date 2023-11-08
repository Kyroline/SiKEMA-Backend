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
		StudentId: input.StudentId,
		Hours:     input.Hours,
	}

	if err := util.DB.Create(&newAbsent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": newAbsent,
	})
}
