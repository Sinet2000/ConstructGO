package dbinit

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeDatabase(client *mongo.Client, ctx context.Context) {
	err := InitializeCategories(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = InitializeItems(client, ctx)
	if err != nil {
		log.Fatal(err)
	}
}
