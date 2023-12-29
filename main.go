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
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	util.Connection("root", "", "127.0.0.1", "3306", "attendance-is")
	// util.Migrate()
	// util.Seed()
	r := RouterSetup()
	r.StaticFS("/files", http.Dir("public"))

	r.POST("/upload", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("file")
		if err != nil {
			// Handle the error, e.g., log it or return an error response to the client
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		// Upload the file to specific dst.
		c.SaveUploadedFile(file, "public/"+file.Filename)

		c.String(http.StatusOK, file.Filename)
	})

	r.Run(":8080")
}

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*", "http://192.168.137.172"},
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
	route.InitSPRoute(util.DB, router)

	return router
}
