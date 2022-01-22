package main

import (
	"StudentManager/database"
	. "StudentManager/middleware"
	"StudentManager/router"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	close := Logs()
	defer close()
	router.SetupRouter()
	ctx := context.Background()
	db := database.SetupDB(ctx)
	defer func(db *mongo.Client, ctx context.Context) {
		if err := db.Disconnect(ctx); err != nil {
			Ln.Fatal(err)
		}
	}(db, ctx)
}
