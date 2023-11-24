package route

import (
	course "attendance-is/controllers/course"
	service "attendance-is/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCourseRoute(db *gorm.DB, r *gin.Engine) {
	courseService := service.NewCourseService(db)
	courseHandler := course.NewCourseHandler(courseService)

	groupLecturer := r.Group("api/lecturer/:lecturerid/course")
	groupLecturer.GET("", courseHandler.GetCourseByLecturer)

	groupPBM := r.Group("api/course")
	groupPBM.GET("", courseHandler.GetCourseByPBM)
}
