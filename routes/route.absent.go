package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	absent "attendance-is/controllers/absent"
	service "attendance-is/services"
)

func InitAbsentRoute(db *gorm.DB, r *gin.Engine) {
	absentService := service.NewAbsentService(db)
	absentHandler := absent.NewAbsentHandler(absentService)

	group := r.Group("/api/absent")
	group.DELETE(":id")
	group.GET("", absentHandler.GetAbsentsByPBM)
	group.GET(":id/excuse", absentHandler.GetExcuse)
	group.GET(":id", absentHandler.GetAbsentByPBM)
	group.PATCH(":id")
	group.POST("")
	group.PUT(":id")

	studentGroup := r.Group("/api/student/:studentid/absent")
	studentGroup.GET("", absentHandler.GetAbsentsByStudent)
	studentGroup.GET(":id", absentHandler.GetAbsentByStudent)
}
