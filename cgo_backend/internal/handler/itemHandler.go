package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sinet2000/cgo_backend/internal/models"
	"github.com/Sinet2000/cgo_backend/internal/repository"
	"github.com/gorilla/mux"
)

type ItemHandler struct {
	repo *repository.ItemRepository
}

func NewItemHandler(r *repository.ItemRepository) *ItemHandler {
	return &ItemHandler{repo: r}
}

// @Summary Create a new item
// @Description Create a new item with the input payload
// @Tags items
// @Accept  json
// @Produce  json
// @Param item body models.Item true "Create item"
// @Success 200 {object} models.Item
// @Router /items [post]
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	result, err := h.repo.CreateItem(r.Context(), item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// @Summary Get all items
// @Description Get all items
// @Tags items
// @Produce  json
// @Success 200 {array} models.Item
// @Router /items [get]
func (h *ItemHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.GetItems(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(items)
}

// @Summary Get an item
// @Description Get an item by id
// @Tags items
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} models.Item
// @Router /items/{id} [get]
func (h *ItemHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	item, err := h.repo.GetItem(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

// @Summary Update an item
// @Description Update an item with the input payload
// @Tags items
// @Accept  json
// @Produce  json
// @Param id path string true "Item ID"
// @Param item body models.Item true "Update item"
// @Success 200 {object} models.Item
// @Router /items/{id} [put]
func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	result, err := h.repo.UpdateItem(r.Context(), id, item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// @Summary Delete an item
// @Description Delete an item by id
// @Tags items
// @Produce  json
// @Param id path string true "Item ID"
// @Success 200 {object} map[string]string "result" "Item deleted"
// @Failure 404 {object} map[string]string "result" "No item found to delete"
// @Router /items/{id} [delete]
func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := h.repo.DeleteItem(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		json.NewEncoder(w).Encode(map[string]string{"result": "No item found to delete"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "Item deleted"})
}
