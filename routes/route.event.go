package route

import (
	event "attendance-is/controllers/event"
	middleware "attendance-is/middlewares"
	service "attendance-is/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitEventRoute(db *gorm.DB, r *gin.Engine) {
	eventService := service.NewEventService(db)
	eventHandler := event.NewEventHandler(eventService)

	//from lecturer
	groupLecturer := r.Group("/api/lecturer/:lecturerid/event", middleware.Auth(), middleware.IsLecturer())
	groupLecturer.DELETE(":id") //can only delete recently created event
	groupLecturer.GET("", eventHandler.GetEvent)
	groupLecturer.GET(":id", eventHandler.FindEvent)
	groupLecturer.PATCH(":id", eventHandler.FinalizeEventHandler)
	groupLecturer.POST("", eventHandler.CreateEventHandler)
	groupLecturer.POST(":id/student", eventHandler.AddStudentToEvent)
	groupLecturer.DELETE(":id/student", eventHandler.RemoveStudent)
	groupLecturer.PUT(":id", eventHandler.FinalizeEventHandler)

	groupStudent := r.Group("/api/student/:studentid/event", middleware.Auth(), middleware.IsStudent())
	groupStudent.GET("")
	groupStudent.GET(":id")
}
