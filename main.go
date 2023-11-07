package main

import (
	route "attendance-is/routes"
	util "attendance-is/utils"
)

func main() {
	util.Connection("root", "", "127.0.0.1", "3306", "attendance-is")
	util.Migrate()
	r := route.Init()
	r.Run()
}
