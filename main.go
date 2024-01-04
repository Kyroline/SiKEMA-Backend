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
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"flag"
	"fmt"
)

func main() {
	isMigrate := flag.Bool("migrate", false, "If set to true, performs database migration before starting the webservice. Optional, use this flag to ensure the required database structure is available before running the application.")
	isSeed := flag.Bool("seed", false, "Seeds the database with initial data before starting the webservice. Optional, use this flag to populate the database with predefined data.")
	host := flag.String("host", "", "Specifies the host address for deploying the webservice. Use this flag to define the desired host address.")
	port := flag.String("port", "8080", "Specifies the port number for deploying the webservice. Use this flag to define the desired host address.")

	flag.Parse()

	dbUser := os.Getenv("ATTENDANCE_DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}

	dbPass := os.Getenv("ATTENDANCE_DB_PASSWORD")

	dbHost := os.Getenv("ATTENDANCE_DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}

	dbPort := os.Getenv("ATTENDANCE_DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}

	dbName := os.Getenv("ATTENDANCE_DB_NAME")
	if dbName == "" {
		dbName = "attendance-is"
	}

	util.Connection(dbUser, dbPass, dbHost, dbPort, dbName)
	r := RouterSetup()
	r.StaticFS("/files", http.Dir("public"))

	r.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {

			c.String(http.StatusBadRequest, err.Error())
			return
		}

		name := strconv.FormatInt(time.Now().UnixMilli(), 10) + ".pdf"

		c.SaveUploadedFile(file, "public/"+name)

		c.JSON(http.StatusOK, gin.H{
			"data": name,
		})
	})

	if *isMigrate {
		util.Migrate()
	}

	if *isSeed {
		util.Seed()
	}
	fmt.Println(*host + ":" + *port)

	r.Run(*host + ":" + *port)
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
