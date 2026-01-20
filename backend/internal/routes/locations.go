package routes

import (
	"backend/internal/utils"
	"encoding/json"
	"os"

	"net/http"
	"path/filepath"
)

func countriesHandler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("locations", "countries.json")
	data, err := os.ReadFile(path)
	if err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "Unable to read countries file", nil)
		return
	}

	var parsed any
	if err := json.Unmarshal(data, &parsed); err != nil {
		utils.JSON(w, http.StatusInternalServerError, false, "Failed to parse countries file", nil)
		return
	}

	utils.JSON(w, http.StatusOK, true, "Countries fetched successfully", parsed)
}

func countryDataHandler(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		utils.JSON(w, http.StatusBadRequest, false, "Missing country code", nil)
		return
	}

	basePath := filepath.Join("locations", code) // e.g., locations/bd

	// Read all three files
	divBytes, err := os.ReadFile(filepath.Join(basePath, "divisions.json"))
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "Division data not found", nil)
		return
	}
	disBytes, err := os.ReadFile(filepath.Join(basePath, "districts.json"))
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "District data not found", nil)
		return
	}
	subBytes, err := os.ReadFile(filepath.Join(basePath, "subdistricts.json"))
	if err != nil {
		utils.JSON(w, http.StatusNotFound, false, "Subdistrict data not found", nil)
		return
	}

	type CountryData struct {
		Code         string          `json:"code"`
		Divisions    json.RawMessage `json:"divisions"`
		Districts    json.RawMessage `json:"districts"`
		SubDistricts json.RawMessage `json:"sub_districts"`
	}

	// Combine using json.RawMessage (no full unmarshal)
	data := CountryData{
		Code:         code,
		Divisions:    divBytes,
		Districts:    disBytes,
		SubDistricts: subBytes,
	}

	// Send response using utils.JSON
	utils.JSON(w, http.StatusOK, true, "Country data fetched", data)
}
