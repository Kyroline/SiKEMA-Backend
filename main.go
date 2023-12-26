//   SiKEMA Api
//    version: 0.1
//    title: SiKEMA Api
//   Schemes: http, https
//   Host: localhost:8000
//   BasePath: /api/v1
//      Consumes:
//      - application/json
//   Produces:
//   - application/json
//   SecurityDefinitions:
//    Bearer:
//     type: apiKey
//     name: Authorization
//     in: header
//   swagger:meta
package main

import (
	route "attendance-is/routes"
	util "attendance-is/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	util.Connection("root", "", "127.0.0.1", "3306", "attendance-is")
	// util.Migrate()
	// util.Seed()
	r := RouterSetup()
	opts := middleware.SwaggerUIOpts{SpecURL: "./swagger.yaml"}
	sh := middleware.SwaggerUI(opts, nil)
	r.GET("/docs", gin.WrapH(sh))
	r.Run(":8080")
}

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	route.InitAbsentRoute(util.DB, router)
	route.InitClassRoute(util.DB, router)
	route.InitEventRoute(util.DB, router)
	route.InitCourseRoute(util.DB, router)
	route.InitStudentRoute(util.DB, router)
	route.InitAuthRoute(util.DB, router)
	route.InitExcuseRoute(util.DB, router)
	route.InitCompensationRoute(util.DB, router)

	return router
}
