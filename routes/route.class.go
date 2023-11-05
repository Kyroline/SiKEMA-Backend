package route

import "github.com/gin-gonic/gin"

func InitClassRoute(r *gin.Engine) {
	group := r.Group("/class")
	group.GET("")
	group.GET(":id")
	group.PATCH(":id")
	group.POST("")
	group.PUT(":id")
}
