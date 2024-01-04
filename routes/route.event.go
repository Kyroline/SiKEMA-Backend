package route

import (
	getEvent "attendance-is/controllers/event/lecturer/get"
	getEventHandler "attendance-is/handlers/event/lecturer/get"

	showEvent "attendance-is/controllers/event/lecturer/show"
	showEventHandler "attendance-is/handlers/event/lecturer/show"

	addStudentEvent "attendance-is/controllers/event/lecturer/add_student"
	addStudentEventHandler "attendance-is/handlers/event/lecturer/add_student"

	removeStudentEvent "attendance-is/controllers/event/lecturer/remove_student"
	removeStudentEventHandler "attendance-is/handlers/event/lecturer/remove_student"

	createEvent "attendance-is/controllers/event/lecturer/create"
	createEventHandler "attendance-is/handlers/event/lecturer/create"

	finalizeEvent "attendance-is/controllers/event/lecturer/finalize"
	finalizeEventHandler "attendance-is/handlers/event/lecturer/finalize"

	middleware "attendance-is/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitEventRoute(db *gorm.DB, r *gin.Engine) {
	getEventRepo := getEvent.NewGetEventRepository(db)
	getEventService := getEvent.NewGetEventService(getEventRepo)
	getEventHandler := getEventHandler.NewGetEventHandler(getEventService)

	showEventRepo := showEvent.NewShowEventRepository(db)
	showEventService := showEvent.NewShowEventService(showEventRepo)
	showEventHandler := showEventHandler.NewShowEventHandler(showEventService)

	addStudentEventRepo := addStudentEvent.NewAddStudentEventRepository(db)
	addStudentEventService := addStudentEvent.NewAddStudentToEventService(addStudentEventRepo)
	addStudentEventHandler := addStudentEventHandler.NewAddStudentEventHandler(addStudentEventService)

	removeStudentEventRepo := removeStudentEvent.NewRemoveStudentEventRepository(db)
	removeStudentEventService := removeStudentEvent.NewRemoveStudentToEventService(removeStudentEventRepo)
	removeStudentEventHandler := removeStudentEventHandler.NewRemoveStudentEventHandler(removeStudentEventService)

	createEventRepo := createEvent.NewCreateEventRepository(db)
	createEventService := createEvent.NewCreateEventService(createEventRepo)
	createEventHandler := createEventHandler.NewCreateEventHandler(createEventService)

	finalizeEventRepo := finalizeEvent.NewFinalizeEventRepository(db)
	finalizeEventService := finalizeEvent.NewFinalizeEventService(finalizeEventRepo)
	finalizeEventHandler := finalizeEventHandler.NewFinalizeEventHandler(finalizeEventService)

	groupLecturer := r.Group("/api/lecturer/:lecturerid/event", middleware.Auth(), middleware.IsLecturer())

	groupLecturer.GET("", getEventHandler.GetEventHandler)
	groupLecturer.GET(":id", showEventHandler.ShowEventHandler)
	groupLecturer.PATCH(":id", finalizeEventHandler.FinalizeEventHandler)
	groupLecturer.POST("", createEventHandler.CreateEventHandler)
	groupLecturer.POST(":id/student", addStudentEventHandler.AddStudentEventHandler)
	groupLecturer.DELETE(":id/student", removeStudentEventHandler.RemoveStudentEvent)

	// groupStudent := r.Group("/api/student/:studentid/event", middleware.Auth(), middleware.IsStudent())
	// groupStudent.GET("")
	// groupStudent.GET(":id")
}
