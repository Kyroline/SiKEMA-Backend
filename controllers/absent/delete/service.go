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

	if err := util.DB.Delete(&model.Absent{}, id).Error; util.HandleGORMError(c, err, (map[bool]int{true: http.StatusNotFound, false: http.StatusInternalServerError})[errors.Is(err, gorm.ErrRecordNotFound)], "") {
		return
	}

	c.Status(http.StatusNoContent)
}
