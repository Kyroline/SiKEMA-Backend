package doPresent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DoPresent(c *gin.Context) {
	var input Input
	c.ShouldBindJSON(&input)

	var student model.Student
	if err := util.DB.Model(model.Student{}).Where("id = ?", input.Nim).Find(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := util.DB.Model(&model.Event{}).Association("Students").Append(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Record successfully created",
	})
}

func MassPresent(c *gin.Context) {
	var input MassInput
	c.ShouldBindJSON(&input)

	var students []model.Student
	if err := util.DB.Model(&[]model.Student{}).Where("id IN ?", input.Nim).Find(&students).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	util.DB.Transaction(func(tx *gorm.DB) error {

		var event model.Event
		if err := tx.Model(&model.Event{}).Where("id = ?", input.EventId).Find(&event).Error; err != nil {
			return err
		}

		if err := tx.Model(&event).Association("Students").Append(&students); err != nil {
			return err
		}
		return nil
	})

	c.JSON(http.StatusCreated, gin.H{
		"message": "Record successfully created",
	})
}
