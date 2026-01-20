package routes

import (
	"backend/internal/middlewares"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middlewares.CtxRole).(string)
	if role == "client" {
		utils.JSON(w, http.StatusForbidden, false, "forbidden", nil)
		return
	}

	type Request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cat := &models.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := models.CreateCategory(ctx, cat); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot create category", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "category created", cat)
}

func updateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middlewares.CtxRole).(string)
	if role == "client" {
		utils.JSON(w, http.StatusForbidden, false, "forbidden", nil)
		return
	}

	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)

	type Request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cat := &models.Category{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := models.UpdateCategory(ctx, cat); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot update category", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "category updated", cat)
}

func deleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middlewares.CtxRole).(string)
	if role == "client" {
		utils.JSON(w, http.StatusForbidden, false, "forbidden", nil)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := models.DeleteCategory(ctx, id); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot delete category", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "category deleted", nil)
}

//

func getCategoriesAndSubcategoriesHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cats, err := models.GetAllCategories(ctx)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch categories", nil)
		return
	}

	scats, err := models.GetAllSubCategories(ctx)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch sub-categories", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "categories fetched", map[string]any{
		"categories":     cats,
		"sub-categories": scats,
	})
}

//

func createSubCategoryHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middlewares.CtxRole).(string)
	if role == "client" {
		utils.JSON(w, http.StatusForbidden, false, "forbidden", nil)
		return
	}

	type Request struct {
		CategoryID  int64  `json:"category_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" || req.CategoryID == 0 {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sc := &models.SubCategory{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := models.CreateSubCategory(ctx, sc); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot create subcategory", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "subcategory created", sc)
}

func updateSubCategoryHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middlewares.CtxRole).(string)
	if role == "client" {
		utils.JSON(w, http.StatusForbidden, false, "forbidden", nil)
		return
	}

	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)

	type Request struct {
		CategoryID  int64  `json:"category_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" || req.CategoryID == 0 {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sc := &models.SubCategory{
		ID:          id,
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := models.UpdateSubCategory(ctx, sc); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot update subcategory", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "subcategory updated", sc)
}

func deleteSubCategoryHandler(w http.ResponseWriter, r *http.Request) {
	role := r.Context().Value(middlewares.CtxRole).(string)
	if role == "client" {
		utils.JSON(w, http.StatusForbidden, false, "forbidden", nil)
		return
	}

	idStr := r.PathValue("id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := models.DeleteSubCategory(ctx, id); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot delete subcategory", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "subcategory deleted", nil)
}
