package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	route "attendance-is/routes"
	util "attendance-is/utils"
)

func main() {
	util.Connection("root", "", "127.0.0.1", "3306", "attendance-is")
	util.Migrate()
	r := SetupRouter()
	r.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	route.InitAbsentRoute(router)
	route.InitClassRoute(router)
	route.InitEventRoute(router)

	return router
}
