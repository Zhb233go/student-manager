package main

import (
	"context"
	"fiber-demo/database"
	"fiber-demo/router"
)

func main() {
	ctx := context.Background()
	database.SetupDB(ctx)
	router.SetupRouter()
}
