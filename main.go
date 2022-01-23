package main

import (
	"StudentManager/database"
	. "StudentManager/middleware"
	"StudentManager/router"
)

func main() {
	database.V8Example()
	close := Logs()
	defer close()
	router.SetupRouter()
}
