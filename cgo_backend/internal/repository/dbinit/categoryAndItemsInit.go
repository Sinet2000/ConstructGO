package dbinit

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"github.com/Sinet2000/cgo_backend/internal/models"
	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/mongo"
)

func generateSampleCategories() []models.Category {
	var categories []models.Category
	for i := 0; i < 3; i++ {
		category := models.Category{
			Name:        gofakeit.Word(),
			Description: gofakeit.Sentence(10),
		}
		categories = append(categories, category)
	}
	return categories
}

func insertCategories(collection *mongo.Collection, ctx context.Context, categories []models.Category) error {
	var categoryInterfaces []interface{}
	for _, category := range categories {
		categoryInterfaces = append(categoryInterfaces, category)
	}

	_, err := collection.InsertMany(ctx, categoryInterfaces)
	if err != nil {
		log.Println("Failed to initialize categories:", err)
		return err
	}

	return nil
}

func InitializeCategoriesAndItems(client *mongo.Client, ctx context.Context, dbName, categoryCollName, itemCollName string) error {
	err := InitializeItems(client, ctx, dbName, itemCollName)
	if err != nil {
		return err
	}

	collection := client.Database(dbName).Collection(categoryCollName)

	// Generate sample categories
	categories := generateSampleCategories()

	// Insert categories into the collection
	err = insertCategories(collection, ctx, categories)
	if err != nil {
		return err
	}

	// Retrieve the items from the "items" collection
	itemsCollection := client.Database(dbName).Collection(itemCollName)
	var items []models.Item
	cursor, err := itemsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Failed to retrieve items:", err)
		return err
	}
	if err := cursor.All(ctx, &items); err != nil {
		log.Println("Failed to decode items:", err)
		return err
	}

	// Add the items to the first created category
	if len(categories) > 0 {
		categories[0].Items = items
	}

	// Update the first created category in the collection
	update := bson.M{
		"$set": bson.M{"items": categories[0].Items},
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": categories[0].ID}, update)
	if err != nil {
		log.Println("Failed to update the first category:", err)
		return err
	}

	log.Println("Initialized categories and items successfully")
	return nil
}

func InitializeItems(client *mongo.Client, ctx context.Context, dbName, itemCollName string) error {
	collection := client.Database(dbName).Collection(itemCollName)

	// Generate sample items using gofakeit
	var items []interface{}
	for i := 0; i < 3; i++ {
		item := models.Item{
			Name:          gofakeit.MinecraftTool(),
			SKU:           gofakeit.UUID(),
			Image:         gofakeit.ImageURL(200, 200),
			Price:         gofakeit.Price(10, 100),
			Weight:        gofakeit.Float64Range(1, 10),
			Quantity:      gofakeit.Number(1, 10),
			ItemType:      gofakeit.RandomString([]string{"PerItem", "PerSize", "PerWeight"}),
			PurchasePrice: gofakeit.Price(5, 50),
			Height:        gofakeit.Float64Range(1, 10),
			Width:         gofakeit.Float64Range(1, 10),
			Length:        gofakeit.Float64Range(1, 10),
		}
		items = append(items, item)
	}

	// Insert items into the collection
	_, err := collection.InsertMany(ctx, items)
	if err != nil {
		log.Println("Failed to initialize items:", err)
		return err
	}

	log.Println("Initialized items successfully")
	return nil
}
