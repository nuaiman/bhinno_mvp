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

func createServiceHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, ok := r.Context().Value(middlewares.CtxUserID).(int64)
	if !ok || userID == 0 {
		utils.JSON(w, http.StatusUnauthorized, false, "unauthorized", nil)
		return
	}

	var req struct {
		CategoryID    int64                  `json:"category_id"`
		SubcategoryID int64                  `json:"subcategory_id"`
		DivisionID    int                    `json:"division_id"`
		DistrictID    int                    `json:"district_id"`
		SubdistrictID int                    `json:"subdistrict_id"`
		Area          string                 `json:"area"`
		Title         string                 `json:"title"`
		Caption       string                 `json:"caption"`
		Description   string                 `json:"description"`
		Price         string                 `json:"price"`
		Features      map[string]interface{} `json:"features"`
		Hours         string                 `json:"hours"`
		Days          []string               `json:"days"`
		PageName      string                 `json:"page_name"`
		PageLink      string                 `json:"page_link"`
		MessengerName string                 `json:"messenger_name"`
		MessengerLink string                 `json:"messenger_link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	service := &models.Service{
		UserID:        userID,
		CategoryID:    req.CategoryID,
		SubcategoryID: req.SubcategoryID,
		DivisionID:    req.DivisionID,
		DistrictID:    req.DistrictID,
		SubdistrictID: req.SubdistrictID,
		Area:          req.Area,
		Title:         req.Title,
		Caption:       req.Caption,
		Description:   req.Description,
		Price:         req.Price,
		Features:      req.Features,
		Hours:         req.Hours,
		Days:          req.Days,
		PageName:      req.PageName,
		PageLink:      req.PageLink,
		MessengerName: req.MessengerName,
		MessengerLink: req.MessengerLink,
	}

	if err := models.CreateService(ctx, service); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot create service", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "service created successfully", map[string]any{
		"service": service,
	})
}

func getServiceHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	idStr := r.PathValue("id")
	serviceID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid service ID", nil)
		return
	}

	service, err := models.GetServiceByID(ctx, serviceID)
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "service not found", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "service fetched", map[string]any{
		"service": service,
	})
}

func updateServiceHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, ok := r.Context().Value(middlewares.CtxUserID).(int64)
	if !ok || userID == 0 {
		utils.JSON(w, http.StatusUnauthorized, false, "unauthorized", nil)
		return
	}

	idStr := r.PathValue("id")
	serviceID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid service ID", nil)
		return
	}

	service, err := models.GetServiceByID(ctx, serviceID)
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "service not found", nil)
		return
	}

	if service.UserID != userID {
		utils.JSON(w, http.StatusForbidden, false, "cannot edit someone else's service", nil)
		return
	}

	var req struct {
		CategoryID    *int64                 `json:"category_id"`
		SubcategoryID *int64                 `json:"subcategory_id"`
		DivisionID    *int                   `json:"division_id"`
		DistrictID    *int                   `json:"district_id"`
		SubdistrictID *int                   `json:"subdistrict_id"`
		Area          *string                `json:"area"`
		Title         *string                `json:"title"`
		Caption       *string                `json:"caption"`
		Description   *string                `json:"description"`
		Price         *string                `json:"price"`
		Features      map[string]interface{} `json:"features"`
		Hours         *string                `json:"hours"`
		Days          []string               `json:"days"`
		PageName      *string                `json:"page_name"`
		PageLink      *string                `json:"page_link"`
		MessengerName *string                `json:"messenger_name"`
		MessengerLink *string                `json:"messenger_link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	if req.CategoryID != nil {
		service.CategoryID = *req.CategoryID
	}
	if req.SubcategoryID != nil {
		service.SubcategoryID = *req.SubcategoryID
	}
	if req.DivisionID != nil {
		service.DivisionID = *req.DivisionID
	}
	if req.DistrictID != nil {
		service.DistrictID = *req.DistrictID
	}
	if req.SubdistrictID != nil {
		service.SubdistrictID = *req.SubdistrictID
	}
	if req.Area != nil {
		service.Area = *req.Area
	}
	if req.Title != nil {
		service.Title = *req.Title
	}
	if req.Caption != nil {
		service.Caption = *req.Caption
	}
	if req.Description != nil {
		service.Description = *req.Description
	}
	if req.Price != nil {
		service.Price = *req.Price
	}
	if req.Features != nil {
		service.Features = req.Features
	}
	if req.Hours != nil {
		service.Hours = *req.Hours
	}
	if req.Days != nil {
		service.Days = req.Days
	}
	if req.PageName != nil {
		service.PageName = *req.PageName
	}
	if req.PageLink != nil {
		service.PageLink = *req.PageLink
	}
	if req.MessengerName != nil {
		service.MessengerName = *req.MessengerName
	}
	if req.MessengerLink != nil {
		service.MessengerLink = *req.MessengerLink
	}

	if err := models.UpdateService(ctx, service); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot update service", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "service updated successfully", map[string]any{
		"service": service,
	})
}

func deleteServiceHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userID, ok := r.Context().Value(middlewares.CtxUserID).(int64)
	if !ok || userID == 0 {
		utils.JSON(w, http.StatusUnauthorized, false, "unauthorized", nil)
		return
	}

	idStr := r.PathValue("id")
	serviceID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid service ID", nil)
		return
	}

	service, err := models.GetServiceByID(ctx, serviceID)
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "service not found", nil)
		return
	}

	if service.UserID != userID {
		utils.JSON(w, http.StatusForbidden, false, "cannot delete someone else's service", nil)
		return
	}

	if err := models.DeleteService(ctx, serviceID); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot delete service", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "service deleted successfully", nil)
}

func filterServicesHandler(w http.ResponseWriter, r *http.Request) {
	divisionIDStr := r.PathValue("division_id")
	districtIDStr := r.PathValue("district_id")
	subdistrictIDStr := r.PathValue("subdistrict_id")
	categoryIDStr := r.PathValue("category_id")
	subcategoryIDStr := r.PathValue("subcategory_id")

	if divisionIDStr == "" || districtIDStr == "" || subdistrictIDStr == "" || categoryIDStr == "" || subcategoryIDStr == "" {
		utils.JSON(w, http.StatusBadRequest, false, "all filter parameters are required", nil)
		return
	}

	divisionID, err := strconv.Atoi(divisionIDStr)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid division_id", nil)
		return
	}

	districtID, err := strconv.Atoi(districtIDStr)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid district_id", nil)
		return
	}

	subdistrictID, err := strconv.Atoi(subdistrictIDStr)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid subdistrict_id", nil)
		return
	}

	categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid category_id", nil)
		return
	}

	subcategoryID, err := strconv.ParseInt(subcategoryIDStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid subcategory_id", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	services, err := models.GetServicesByFilters(ctx, divisionID, districtID, subdistrictID, categoryID, subcategoryID)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch services", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "services fetched successfully", map[string]any{
		"services": services,
	})
}
