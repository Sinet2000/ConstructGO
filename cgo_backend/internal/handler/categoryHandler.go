package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Sinet2000/cgo_backend/internal/models"
	"github.com/Sinet2000/cgo_backend/internal/repository"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	repo repository.CategoryRepositoryInterface
}

func NewCategoryHandler(r repository.CategoryRepositoryInterface) *CategoryHandler {
	return &CategoryHandler{repo: r}
}

// @Summary Create a new category
// @Description Create a new category with the input payload
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body models.Category true "Create category"
// @Success 200 {object} models.Category
// @Router /categories [post]
func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	result, err := h.repo.CreateCategory(r.Context(), category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Produce  json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetCategories(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

// @Summary Get a category
// @Description Get a category by id
// @Tags categories
// @Produce  json
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	category, err := h.repo.GetCategory(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}

// @Summary Update a category
// @Description Update a category with the input payload
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path string true "Category ID"
// @Param category body models.Category true "Update category"
// @Success 200 {object} models.Category
// @Router /categories/{id} [put]
func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)

	result, err := h.repo.UpdateCategory(r.Context(), id, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// @Summary Delete a category
// @Description Delete a category by id
// @Tags categories
// @Produce  json
// @Param id path string true "Category ID"
// @Success 200 {object} map[string]string "result" "Category deleted"
// @Failure 404 {object} map[string]string "result" "No category found to delete"
// @Router /categories/{id} [delete]
func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	result, err := h.repo.DeleteCategory(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		json.NewEncoder(w).Encode(map[string]string{"result": "No category found to delete"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "Category deleted"})
}

// @Summary Add a subcategory to a category
// @Description Add a subcategory to a category with the input payload
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path string true "Category ID"
// @Param subcategory body models.SubCategory true "Add subcategory"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]string "result" "Failed to add subcategory"
// @Router /categories/{id}/subcategories [post]
func (h *CategoryHandler) AddSubCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var subcategory models.SubCategory
	err := json.NewDecoder(r.Body).Decode(&subcategory)
	if err != nil {
		http.Error(w, "Invalid subcategory data", http.StatusBadRequest)
		return
	}

	err = h.repo.AddSubCategory(r.Context(), id, subcategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"result": "Subcategory added successfully"})
}
