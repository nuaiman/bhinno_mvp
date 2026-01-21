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
		CountryCode           string                 `json:"country_code"`
		CategoryID            int64                  `json:"category_id"`
		SubcategoryID         int64                  `json:"subcategory_id"`
		StateID               int                    `json:"state_id"`
		AdministrativeAreaID  int                    `json:"administrative_area_id"`
		SubAdministrativeArea int                    `json:"sub_administrative_area_id"`
		Area                  string                 `json:"area"`
		Title                 string                 `json:"title"`
		Caption               string                 `json:"caption"`
		Description           string                 `json:"description"`
		Price                 string                 `json:"price"`
		Features              map[string]interface{} `json:"features"`
		Hours                 string                 `json:"hours"`
		Days                  []string               `json:"days"`
		PageName              string                 `json:"page_name"`
		PageLink              string                 `json:"page_link"`
		MessengerName         string                 `json:"messenger_name"`
		MessengerLink         string                 `json:"messenger_link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	service := &models.Service{
		UserID:                  userID,
		CountryCode:             req.CountryCode,
		CategoryID:              req.CategoryID,
		SubcategoryID:           req.SubcategoryID,
		StateID:                 req.StateID,
		AdministrativeAreaID:    req.AdministrativeAreaID,
		SubAdministrativeAreaID: req.SubAdministrativeArea,
		Area:                    req.Area,
		Title:                   req.Title,
		Caption:                 req.Caption,
		Description:             req.Description,
		Price:                   req.Price,
		Features:                req.Features,
		Hours:                   req.Hours,
		Days:                    req.Days,
		PageName:                req.PageName,
		PageLink:                req.PageLink,
		MessengerName:           req.MessengerName,
		MessengerLink:           req.MessengerLink,
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
		CountryCode             *string                `json:"country_code"`
		CategoryID              *int64                 `json:"category_id"`
		SubcategoryID           *int64                 `json:"subcategory_id"`
		StateID                 *int                   `json:"state_id"`
		AdministrativeAreaID    *int                   `json:"administrative_area_id"`
		SubAdministrativeAreaID *int                   `json:"sub_administrative_area_id"`
		Area                    *string                `json:"area"`
		Title                   *string                `json:"title"`
		Caption                 *string                `json:"caption"`
		Description             *string                `json:"description"`
		Price                   *string                `json:"price"`
		Features                map[string]interface{} `json:"features"`
		Hours                   *string                `json:"hours"`
		Days                    []string               `json:"days"`
		PageName                *string                `json:"page_name"`
		PageLink                *string                `json:"page_link"`
		MessengerName           *string                `json:"messenger_name"`
		MessengerLink           *string                `json:"messenger_link"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	if req.CountryCode != nil {
		service.CountryCode = *req.CountryCode
	}
	if req.CategoryID != nil {
		service.CategoryID = *req.CategoryID
	}
	if req.SubcategoryID != nil {
		service.SubcategoryID = *req.SubcategoryID
	}
	if req.StateID != nil {
		service.StateID = *req.StateID
	}
	if req.AdministrativeAreaID != nil {
		service.AdministrativeAreaID = *req.AdministrativeAreaID
	}
	if req.SubAdministrativeAreaID != nil {
		service.SubAdministrativeAreaID = *req.SubAdministrativeAreaID
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

func getFilteredServicesHandler(w http.ResponseWriter, r *http.Request) {
	country := r.PathValue("country_code")
	stateStr := r.PathValue("state_id")
	adminStr := r.PathValue("administrative_area_id")
	subadminStr := r.PathValue("sub_administrative_area_id")
	categoryStr := r.PathValue("category_id")
	subcategoryStr := r.PathValue("subcategory_id")

	if country == "" || stateStr == "" || adminStr == "" || subadminStr == "" || categoryStr == "" || subcategoryStr == "" {
		utils.JSON(w, http.StatusBadRequest, false, "all filter parameters are required", nil)
		return
	}

	stateID, err := strconv.Atoi(stateStr)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid state_id", nil)
		return
	}
	adminID, err := strconv.Atoi(adminStr)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid administrative_area_id", nil)
		return
	}
	subadminID, err := strconv.Atoi(subadminStr)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid sub_administrative_area_id", nil)
		return
	}
	categoryID, err := strconv.ParseInt(categoryStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid category_id", nil)
		return
	}
	subcategoryID, err := strconv.ParseInt(subcategoryStr, 10, 64)
	if err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid subcategory_id", nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	services, err := models.GetServicesByFilters(ctx, country, stateID, adminID, subadminID, categoryID, subcategoryID)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch services", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "services fetched successfully", map[string]any{
		"services": services,
	})
}
