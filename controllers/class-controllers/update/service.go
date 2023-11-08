package updateClass

import (
	model "attendance-is/models"
	util "attendance-is/utils"

	"github.com/gin-gonic/gin"
)

func UpdateClass(c *gin.Context) {
	id := c.Param("id")
	var class model.Class
	util.DB.Where("id = ?", id).Find(&class)

	var input Input
	c.ShouldBindJSON(&input)

	class.Name = input.Name

	util.DB.Save(&class)
}

func AppendStudent(c *gin.Context) {
	id := c.Param("id")
	var class model.Class
	util.DB.Where("id = ?", id).Find(&class)

	var input AppendInput
	c.ShouldBindJSON(&input)

	var students []model.Student
	util.DB.Where("id IN ?", input.Students).Find(&students)

	util.DB.Model(&class).Association("Students").Append(students)
}

func ReplaceStudent(c *gin.Context) {
	id := c.Param("id")
	var class model.Class
	util.DB.Where("id = ?", id).Find(&class)

	var input AppendInput
	c.ShouldBindJSON(&input)

	var students []model.Student
	util.DB.Where("id IN ?", input.Students).Find(&students)

	util.DB.Model(&class).Association("Students").Replace(students)
}

func RemoveStudent(c *gin.Context) {
	id := c.Param("id")
	var class model.Class
	util.DB.Where("id = ?", id).Find(&class)

	var input AppendInput
	c.ShouldBindJSON(&input)

	var students []model.Student
	util.DB.Where("id IN ?", input.Students).Find(&students)

	util.DB.Model(&class).Association("Students").Delete(students)
}
