package routes

import (
	"backend/internal/middlewares"
	"backend/internal/models"
	"backend/internal/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// List all countries (for users) – without JSON fields
func getCountriesHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	countries, err := models.GetAllCountries(ctx)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot fetch countries", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "countries fetched", map[string]any{
		"countries": countries,
	})
}

// Get full country details by code
func getCountryHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	code := r.PathValue("code")
	if code == "" {
		utils.JSON(w, http.StatusBadRequest, false, "country code required", nil)
		return
	}

	country, err := models.GetLocationByCode(ctx, code)
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "country not found", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "country fetched", map[string]any{
		"country": country,
	})
}

// Admin: create a new country
func createLocationHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userRole, ok := r.Context().Value(middlewares.CtxRole).(string)
	if !ok || (userRole != middlewares.CtxRoleSuperAdmin && userRole != middlewares.CtxRoleAdmin) {
		utils.JSON(w, http.StatusForbidden, false, "admin access required", nil)
		return
	}

	var req models.Location
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	if err := models.CreateLocation(ctx, &req); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot create country", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "country created", map[string]any{
		"country": req,
	})
}

// Admin: update country
func updateLocationHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userRole, ok := r.Context().Value(middlewares.CtxRole).(string)
	if !ok || (userRole != middlewares.CtxRoleSuperAdmin && userRole != middlewares.CtxRoleAdmin) {
		utils.JSON(w, http.StatusForbidden, false, "admin access required", nil)
		return
	}

	code := r.PathValue("code")
	if code == "" {
		utils.JSON(w, http.StatusBadRequest, false, "country code required", nil)
		return
	}

	country, err := models.GetLocationByCode(ctx, code)
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "country not found", nil)
		return
	}

	var req models.Location
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSON(w, http.StatusBadRequest, false, "invalid request body", nil)
		return
	}

	// Merge fields
	if req.CountryName != "" {
		country.CountryName = req.CountryName
	}
	if req.CountryFlag != "" {
		country.CountryFlag = req.CountryFlag
	}
	if req.States != nil {
		country.States = req.States
	}
	if req.AdministrativeAreas != nil {
		country.AdministrativeAreas = req.AdministrativeAreas
	}
	if req.SubAdministrativeAreas != nil {
		country.SubAdministrativeAreas = req.SubAdministrativeAreas
	}

	if err := models.UpdateLocation(ctx, country); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot update country", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "country updated", map[string]any{
		"country": country,
	})
}

// Admin: delete country
func deleteLocationHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	userRole, ok := r.Context().Value(middlewares.CtxRole).(string)
	if !ok || (userRole != middlewares.CtxRoleSuperAdmin && userRole != middlewares.CtxRoleAdmin) {
		utils.JSON(w, http.StatusForbidden, false, "admin access required", nil)
		return
	}

	code := r.PathValue("code")
	if code == "" {
		utils.JSON(w, http.StatusBadRequest, false, "country code required", nil)
		return
	}

	if err := models.DeleteLocation(ctx, code); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "cannot delete country", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "country deleted", nil)
}
