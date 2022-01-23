package main

import (
	. "StudentManager/middleware"
	"StudentManager/router"
)

func main() {
	close := Logs()
	defer close()
	router.SetupRouter()
}
