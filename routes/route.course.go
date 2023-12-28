package route

import (
	course "attendance-is/controllers/course"
	middleware "attendance-is/middlewares"

	getCourseStudent "attendance-is/controllers/course/student/get"
	getAllCourseStudent "attendance-is/controllers/course/student/getAll"
	getCourseHandler "attendance-is/handlers/course/student/get"
	getAllCourseHandler "attendance-is/handlers/course/student/getAll"
	service "attendance-is/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCourseRoute(db *gorm.DB, r *gin.Engine) {

	getCourseRepository := getCourseStudent.NewGetCourseRepository(db)
	getCourseService := getCourseStudent.NewGetCourseService(getCourseRepository)
	getCourseHandler := getCourseHandler.NewGetCourseHandler(getCourseService)

	getAllCourseRepository := getAllCourseStudent.NewGetAllCourseRepository(db)
	getAllCourseService := getAllCourseStudent.NewGetAllCourseService(getAllCourseRepository)
	getAllCourseHandler := getAllCourseHandler.NewGetAllCourseHandler(getAllCourseService)

	courseService := service.NewCourseService(db)
	courseHandler := course.NewCourseHandler(courseService)

	groupStudent := r.Group("api/student/:studentid/course", middleware.Auth(), middleware.IsStudent())
	groupStudent.GET("", getAllCourseHandler.GetAllCourseHandler)
	groupStudent.GET(":id", getCourseHandler.GetCourseHandler)

	groupLecturer := r.Group("api/lecturer/:lecturerid/course", middleware.Auth(), middleware.IsLecturer())
	groupLecturer.GET("", courseHandler.GetCourseByLecturer)

	groupPBM := r.Group("api/course")
	groupPBM.GET("", courseHandler.GetCourseByPBM)
}
