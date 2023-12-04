package main

import (
	route "attendance-is/routes"
	util "attendance-is/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	util.Connection("root", "", "127.0.0.1", "3306", "attendance-is")
	util.Migrate()
	// util.Seed()
	r := RouterSetup()
	r.Run(":8000")
}

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	// route.InitAbsentRoute(util.DB, router)
	// route.InitClassRoute(util.DB, router)
	route.InitEventRoute(util.DB, router)
	route.InitCourseRoute(util.DB, router)
	// route.InitStudentRoute(util.DB, router)
	// route.InitExcuseRoute(util.DB, router)

	return router
}
