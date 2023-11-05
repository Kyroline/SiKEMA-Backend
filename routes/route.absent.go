package route

import "github.com/gin-gonic/gin"

func InitAbsentRoute(r *gin.Engine) {
	group := r.Group("/absent")
	group.DELETE(":id")
	group.GET("")
	group.GET(":id")
	group.PATCH(":id")
	group.POST("")
	group.PUT(":id")
}
