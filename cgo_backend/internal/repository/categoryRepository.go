package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Sinet2000/cgo_backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepositoryInterface interface {
	CreateCategory(ctx context.Context, category models.Category) (*mongo.InsertOneResult, error)
	GetCategory(ctx context.Context, id string) (models.Category, error)
	GetCategories(ctx context.Context) ([]models.Category, error)
	UpdateCategory(ctx context.Context, id string, category models.Category) (*mongo.UpdateResult, error)
	DeleteCategory(ctx context.Context, id string) (*mongo.DeleteResult, error)
	AddSubCategory(ctx context.Context, categoryID string, sub models.Category) error
}

type CategoryRepository struct {
	db             *mongo.Client
	DatabaseName   string
	CollectionName string
}

func NewCategoryRepository(db *mongo.Client, databaseName, collectionName string) *CategoryRepository {
	return &CategoryRepository{db: db, DatabaseName: databaseName, CollectionName: collectionName}
}

// Create a new category
func (r *CategoryRepository) CreateCategory(ctx context.Context, category models.Category) (*mongo.InsertOneResult, error) {
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)
	result, err := collection.InsertOne(ctx, category)
	return result, err
}

// Get a category by ID
func (r *CategoryRepository) GetCategory(ctx context.Context, id string) (models.Category, error) {
	var category models.Category
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)

	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return category, fmt.Errorf("failed to get category with ID %s: %w", id, err)
	}

	return category, err
}

// Get all categories
func (r *CategoryRepository) GetCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var category models.Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}
	return categories, nil
}

// Update a category
func (r *CategoryRepository) UpdateCategory(ctx context.Context, id string, category models.Category) (*mongo.UpdateResult, error) {
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": category})
	if err != nil {
		return nil, fmt.Errorf("failed to update category with ID %s: %w", id, err)
	}

	return result, err
}

// Delete a category
func (r *CategoryRepository) DeleteCategory(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	collection := r.db.Database(r.DatabaseName).Collection(r.CollectionName)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, fmt.Errorf("failed to delete category with ID %s: %w", id, err)
	}

	return result, err
}

// Add a subcategory to a category
func (r *CategoryRepository) AddSubCategory(ctx context.Context, categoryID string, sub models.Category) error {
	category, err := r.GetCategory(ctx, categoryID)
	if err != nil {
		return err
	}

	if category.HasChildCategory(sub) {
		return errors.New("subcategory already exists in this category")
	}

	category.AddChildCategory(sub)
	_, err = r.UpdateCategory(ctx, categoryID, category)
	if err != nil {
		return fmt.Errorf("failed to add subcategory to category with ID %s: %w", categoryID, err)
	}

	return err
}
