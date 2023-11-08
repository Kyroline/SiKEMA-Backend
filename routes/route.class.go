package route

import (
	"github.com/gin-gonic/gin"

	updateClass "attendance-is/controllers/class-controllers/update"
)

func InitClassRoute(r *gin.Engine) {
	group := r.Group("/api/class")
	group.DELETE(":id/student", updateClass.RemoveStudent)
	group.GET("")
	group.GET(":id")
	group.PATCH(":id", updateClass.UpdateClass)
	group.PATCH(":id/student", updateClass.ReplaceStudent)
	group.POST("")
	group.POST(":id/student", updateClass.AppendStudent)
	group.PUT(":id", updateClass.UpdateClass)
	group.PUT(":id/student", updateClass.ReplaceStudent)
}
