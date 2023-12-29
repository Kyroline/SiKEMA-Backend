package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	absent "attendance-is/controllers/absent"
	middleware "attendance-is/middlewares"
	service "attendance-is/services"
)

func InitAbsentRoute(db *gorm.DB, r *gin.Engine) {
	absentService := service.NewAbsentService(db)
	absentHandler := absent.NewAbsentHandler(absentService)

	group := r.Group("/api/pbm/absent", middleware.Auth(), middleware.IsPBM())
	group.GET("", absentHandler.GetAbsentsByPBM)
	group.GET(":id/excuse", absentHandler.GetExcuse)
	group.GET(":id", absentHandler.GetAbsentByPBM)

	studentGroup := r.Group("/api/student/:studentid/absent", middleware.Auth(), middleware.IsStudent())
	studentGroup.GET("", absentHandler.GetAbsentsByStudent)
	studentGroup.GET(":id", absentHandler.GetAbsentByStudent)
}
