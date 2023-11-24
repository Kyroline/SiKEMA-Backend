package route

import (
	event "attendance-is/controllers/event"
	service "attendance-is/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitEventRoute(db *gorm.DB, r *gin.Engine) {
	eventService := service.NewEventService(db)
	eventHandler := event.NewEventHandler(eventService)

	//from lecturer
	groupLecturer := r.Group("/api/lecturer/:lecturerid/event")
	groupLecturer.DELETE(":id") //can only delete recently created event
	groupLecturer.GET("", eventHandler.GetEvent)
	groupLecturer.GET(":id", eventHandler.FindEvent)
	groupLecturer.PATCH(":id", eventHandler.FinalizeEventHandler)
	groupLecturer.POST("", eventHandler.CreateEventHandler)
	groupLecturer.POST(":id/students", eventHandler.AddStudentToEvent)
	groupLecturer.PUT(":id", eventHandler.FinalizeEventHandler)
}
