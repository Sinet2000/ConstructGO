package main

import (
	"log"
	"net/http"

	handlers "github.com/Sinet2000/cgo_backend/internal/handler"
	"github.com/Sinet2000/cgo_backend/internal/repository"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	DBName           = "inventory"
	CategoryCollName = "categories"
	ItemCollName     = "items"
	ConnectionString = "mongodb://root:secret@localhost:27017"
)

func main() {
	client, ctx, cancel := repository.ConnectDB(ConnectionString, DBName, CategoryCollName, ItemCollName)
	defer cancel()
	defer client.Disconnect(ctx)

	categoryRepo := repository.NewCategoryRepository(client, DBName, CategoryCollName)
	itemRepo := repository.NewItemRepository(client, DBName, ItemCollName)

	categoryHandler := handlers.NewCategoryHandler(categoryRepo)
	itemHandler := handlers.NewItemHandler(itemRepo)

	router := mux.NewRouter()

	router.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories", categoryHandler.GetCategories).Methods("GET")
	router.HandleFunc("/categories/{id}", categoryHandler.GetCategory).Methods("GET")
	router.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/items", itemHandler.CreateItem).Methods("POST")
	router.HandleFunc("/items", itemHandler.GetItems).Methods("GET")
	router.HandleFunc("/items/{id}", itemHandler.GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", itemHandler.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", itemHandler.DeleteItem).Methods("DELETE")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
