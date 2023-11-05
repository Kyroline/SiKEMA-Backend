package route

import (
	createEvent "attendance-is/controllers/event-controllers/create"
	doPresent "attendance-is/controllers/event-controllers/do-present"
	updateEvent "attendance-is/controllers/event-controllers/update"

	"github.com/gin-gonic/gin"
)

func InitEventRoute(r *gin.Engine) {
	group := r.Group("/event")
	group.DELETE(":id") //can only delete recently created event
	group.GET("")       //
	group.GET(":id")
	group.PATCH(":id", updateEvent.UpdateEvent)
	group.POST("", createEvent.CreateEvent)
	group.POST(":id/record", doPresent.DoPresent)
	group.POST(":id/massrecord", doPresent.MassPresent)
	group.PUT(":id", updateEvent.UpdateEvent)
}
