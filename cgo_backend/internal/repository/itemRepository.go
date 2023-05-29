package repository

import (
	"context"
	"errors"

	"github.com/Sinet2000/cgo_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ItemRepository struct {
	db             *mongo.Client
	DatabaseName   string
	CollectionName string
}

// Constructor function for ItemRepository
func NewItemRepository(db *mongo.Client, databaseName, collectionName string) *ItemRepository {
	return &ItemRepository{db: db, DatabaseName: databaseName, CollectionName: collectionName}
}

// Create a new item
// Create a new item
func (r *ItemRepository) CreateItem(ctx context.Context, item models.Item) (*mongo.InsertOneResult, error) {
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)
	result, err := collection.InsertOne(ctx, item)
	if err != nil {
		return nil, errors.New("unable to create item: " + err.Error())
	}

	return result, nil
}

// Get a item by ID
func (r *ItemRepository) GetItem(ctx context.Context, id string) (models.Item, error) {
	var item models.Item
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return models.Item{}, errors.New("unable to fetch item: " + err.Error())
	}

	return item, nil
}

// Get all items
func (r *ItemRepository) GetItems(ctx context.Context) ([]models.Item, error) {
	var items []models.Item
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("unable to fetch items: " + err.Error())
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var item models.Item
		cursor.Decode(&item)
		items = append(items, item)
	}

	return items, nil
}

// Update a item
func (r *ItemRepository) UpdateItem(ctx context.Context, id string, item models.Item) (*mongo.UpdateResult, error) {
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": item})
	if err != nil {
		return nil, errors.New("unable to update item: " + err.Error())
	}

	return result, nil
}

// Delete a item
func (r *ItemRepository) DeleteItem(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, errors.New("unable to delete item: " + err.Error())
	}

	return result, nil
}
