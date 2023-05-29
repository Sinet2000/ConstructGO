package repository

import (
	"context"
	"github.com/Sinet2000/cgo_backend/internal/repository/dbinit"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(connString, dbName, categoryCollName, itemCollName string) (*mongo.Client, context.Context, context.CancelFunc) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the database exists
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var dbExists bool
	for _, dbName := range databases {
		if dbName == dbName {
			dbExists = true
			break
		}
	}

	if !dbExists {
		dbinit.InitializeCategoriesAndItems(client, ctx, dbName, categoryCollName, itemCollName)
	}

	return client, ctx, cancel
}

//func initializeDatabase(client *mongo.Client, ctx context.Context) {
//	// Here, you can initialize your database, create collections and seed initial data.
//
//	db := client.Database("inventory")
//
//	// Create "categories" collection and add initial categories
//	categories, err := db.Collection("categories").InsertMany(ctx, []interface{}{
//		models.Category{ID: "1", Name: "Category1", Description: "This is category 1"},
//		models.Category{ID: "2", Name: "Category2", Description: "This is category 2"},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Create "items" collection and add initial items
//	// Let's assume every item has a reference to a category. Use the IDs of the created categories.
//	items, err := db.Collection("items").InsertMany(ctx, []interface{}{
//		models.Item{ID: "1", Name: "Item1", Description: "This is item 1", CategoryID: categories.InsertedIDs[0].(string)},
//		models.Item{ID: "2", Name: "Item2", Description: "This is item 2", CategoryID: categories.InsertedIDs[1].(string)},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Log IDs of inserted data
//	log.Printf("Inserted categories with IDs: %v", categories.InsertedIDs)
//	log.Printf("Inserted items with IDs: %v", items.InsertedIDs)
//}
