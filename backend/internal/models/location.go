package models

import (
	"backend/internal/db"
	"context"
	"time"
)

type Location struct {
	CountryCode            string                 `json:"country_code"`
	CountryName            string                 `json:"country_name"`
	CountryFlag            string                 `json:"country_flag"`
	States                 map[string]interface{} `json:"states,omitempty"`
	AdministrativeAreas    map[string]interface{} `json:"administrative_areas,omitempty"`
	SubAdministrativeAreas map[string]interface{} `json:"sub_administrative_areas,omitempty"`
	CreatedAt              time.Time              `json:"created_at"`
}

// Get all countries (for users, without JSON fields)
func GetAllCountries(ctx context.Context) ([]*Location, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT country_code, country_name, country_flag
		FROM locations
		ORDER BY country_name
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []*Location
	for rows.Next() {
		loc := &Location{}
		if err := rows.Scan(&loc.CountryCode, &loc.CountryName, &loc.CountryFlag); err != nil {
			return nil, err
		}
		locations = append(locations, loc)
	}
	return locations, nil
}

// Get full country by code
func GetLocationByCode(ctx context.Context, code string) (*Location, error) {
	loc := &Location{}
	err := db.Pool.QueryRow(ctx, `
		SELECT country_code, country_name, country_flag, states, administrative_areas, sub_administrative_areas, created_at
		FROM locations
		WHERE country_code=$1
	`, code).Scan(
		&loc.CountryCode, &loc.CountryName, &loc.CountryFlag,
		&loc.States, &loc.AdministrativeAreas, &loc.SubAdministrativeAreas,
		&loc.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return loc, nil
}

// Admin: create location
func CreateLocation(ctx context.Context, loc *Location) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO locations
		(country_code, country_name, country_flag, states, administrative_areas, sub_administrative_areas)
		VALUES ($1,$2,$3,$4,$5,$6)
	`, loc.CountryCode, loc.CountryName, loc.CountryFlag, loc.States, loc.AdministrativeAreas, loc.SubAdministrativeAreas)
	return err
}

// Admin: update location
func UpdateLocation(ctx context.Context, loc *Location) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE locations
		SET country_name=$1, country_flag=$2, states=$3, administrative_areas=$4, sub_administrative_areas=$5
		WHERE country_code=$6
	`, loc.CountryName, loc.CountryFlag, loc.States, loc.AdministrativeAreas, loc.SubAdministrativeAreas, loc.CountryCode)
	return err
}

// Admin: delete location
func DeleteLocation(ctx context.Context, code string) error {
	_, err := db.Pool.Exec(ctx, `DELETE FROM locations WHERE country_code=$1`, code)
	return err
}
