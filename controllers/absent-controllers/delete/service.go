package deleteAbsent

import (
	model "attendance-is/models"
	util "attendance-is/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteAbsent(c *gin.Context) {
	id := c.Param("id")

	if err := util.DB.Delete(&model.Absent{}, id).Error; err != nil {
		code := http.StatusInternalServerError

		code = (map[bool]int{true: http.StatusNotFound, false: code})[errors.Is(err, gorm.ErrRecordNotFound)]

		c.JSON(code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
