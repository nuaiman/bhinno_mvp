package models

import (
	"backend/internal/db"
	"context"
	"time"
)

type Service struct {
	ID                      int64                  `json:"id"`
	Active                  bool                   `json:"active"`
	UserID                  int64                  `json:"user_id"`
	CountryCode             string                 `json:"country_code"`
	CategoryID              int64                  `json:"category_id"`
	SubcategoryID           int64                  `json:"subcategory_id"`
	StateID                 int                    `json:"state_id"`
	AdministrativeAreaID    int                    `json:"administrative_area_id"`
	SubAdministrativeAreaID int                    `json:"sub_administrative_area_id"`
	Area                    string                 `json:"area"`
	Title                   string                 `json:"title"`
	Caption                 string                 `json:"caption"`
	Description             string                 `json:"description"`
	Price                   string                 `json:"price"`
	Features                map[string]interface{} `json:"features,omitempty"`
	Hours                   string                 `json:"hours,omitempty"`
	Days                    []string               `json:"days"`
	PageName                string                 `json:"page_name,omitempty"`
	PageLink                string                 `json:"page_link,omitempty"`
	MessengerName           string                 `json:"messenger_name,omitempty"`
	MessengerLink           string                 `json:"messenger_link,omitempty"`
	CreatedAt               time.Time              `json:"created_at"`
}

func CreateService(ctx context.Context, s *Service) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO services (
			active, user_id, country_code, category_id, subcategory_id,
			state_id, administrative_area_id, sub_administrative_area_id,
			area, title, caption, description, price,
			features, hours, days,
			page_name, page_link, messenger_name, messenger_link
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20
		)
	`, s.Active, s.UserID, s.CountryCode, s.CategoryID, s.SubcategoryID,
		s.StateID, s.AdministrativeAreaID, s.SubAdministrativeAreaID,
		s.Area, s.Title, s.Caption, s.Description, s.Price,
		s.Features, s.Hours, s.Days,
		s.PageName, s.PageLink, s.MessengerName, s.MessengerLink,
	)
	return err
}

func GetServiceByID(ctx context.Context, id int64) (*Service, error) {
	s := &Service{}
	err := db.Pool.QueryRow(ctx, `
		SELECT
			id, active, user_id, country_code, category_id, subcategory_id,
			state_id, administrative_area_id, sub_administrative_area_id,
			area, title, caption, description, price,
			features, hours, days,
			page_name, page_link, messenger_name, messenger_link,
			created_at
		FROM services
		WHERE id=$1
	`, id).Scan(
		&s.ID, &s.Active, &s.UserID, &s.CountryCode, &s.CategoryID, &s.SubcategoryID,
		&s.StateID, &s.AdministrativeAreaID, &s.SubAdministrativeAreaID,
		&s.Area, &s.Title, &s.Caption, &s.Description, &s.Price,
		&s.Features, &s.Hours, &s.Days,
		&s.PageName, &s.PageLink, &s.MessengerName, &s.MessengerLink,
		&s.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func GetServicesByFilters(ctx context.Context, country string, stateID, adminID, subadminID int, categoryID, subcategoryID int64) ([]*Service, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT 
			id, active, user_id, country_code, category_id, subcategory_id,
			state_id, administrative_area_id, sub_administrative_area_id,
			area, title, caption, description, price,
			features, hours, days,
			page_name, page_link, messenger_name, messenger_link,
			created_at
		FROM services
		WHERE country_code=$1 AND state_id=$2 AND administrative_area_id=$3 AND sub_administrative_area_id=$4
		  AND category_id=$5 AND subcategory_id=$6
		  AND active=TRUE
		ORDER BY created_at DESC
	`, country, stateID, adminID, subadminID, categoryID, subcategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []*Service
	for rows.Next() {
		s := &Service{}
		if err := rows.Scan(
			&s.ID, &s.Active, &s.UserID, &s.CountryCode, &s.CategoryID, &s.SubcategoryID,
			&s.StateID, &s.AdministrativeAreaID, &s.SubAdministrativeAreaID,
			&s.Area, &s.Title, &s.Caption, &s.Description, &s.Price,
			&s.Features, &s.Hours, &s.Days,
			&s.PageName, &s.PageLink, &s.MessengerName, &s.MessengerLink,
			&s.CreatedAt,
		); err != nil {
			return nil, err
		}
		services = append(services, s)
	}

	return services, nil
}

func UpdateService(ctx context.Context, s *Service) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE services
		SET active=$1, country_code=$2, category_id=$3, subcategory_id=$4,
		    state_id=$5, administrative_area_id=$6, sub_administrative_area_id=$7,
		    area=$8, title=$9, caption=$10, description=$11,
		    price=$12, features=$13, hours=$14, days=$15,
		    page_name=$16, page_link=$17, messenger_name=$18, messenger_link=$19
		WHERE id=$20
	`, s.Active, s.CountryCode, s.CategoryID, s.SubcategoryID,
		s.StateID, s.AdministrativeAreaID, s.SubAdministrativeAreaID,
		s.Area, s.Title, s.Caption, s.Description,
		s.Price, s.Features, s.Hours, s.Days,
		s.PageName, s.PageLink, s.MessengerName, s.MessengerLink,
		s.ID,
	)
	return err
}

func DeleteService(ctx context.Context, id int64) error {
	_, err := db.Pool.Exec(ctx, `DELETE FROM services WHERE id=$1`, id)
	return err
}

func ListServices(ctx context.Context) ([]*Service, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT 
			id, active, user_id, country_code, category_id, subcategory_id,
			state_id, administrative_area_id, sub_administrative_area_id,
			area, title, caption, description, price,
			features, hours, days,
			page_name, page_link, messenger_name, messenger_link,
			created_at
		FROM services
		WHERE active=TRUE
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []*Service
	for rows.Next() {
		s := &Service{}
		if err := rows.Scan(
			&s.ID, &s.Active, &s.UserID, &s.CountryCode, &s.CategoryID, &s.SubcategoryID,
			&s.StateID, &s.AdministrativeAreaID, &s.SubAdministrativeAreaID,
			&s.Area, &s.Title, &s.Caption, &s.Description, &s.Price,
			&s.Features, &s.Hours, &s.Days,
			&s.PageName, &s.PageLink, &s.MessengerName, &s.MessengerLink,
			&s.CreatedAt,
		); err != nil {
			return nil, err
		}
		services = append(services, s)
	}

	return services, nil
}
