package models

import (
	"backend/internal/db"
	"context"
	"time"
)

type Service struct {
	ID            int64                  `json:"id"`
	Active        bool                   `json:"active"`
	UserID        int64                  `json:"user_id"`
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
	Features      map[string]interface{} `json:"features,omitempty"`
	Hours         string                 `json:"hours,omitempty"`
	Days          []string               `json:"days"`
	PageName      string                 `json:"page_name,omitempty"`
	PageLink      string                 `json:"page_link,omitempty"`
	MessengerName string                 `json:"messenger_name,omitempty"`
	MessengerLink string                 `json:"messenger_link,omitempty"`
	CreatedAt     time.Time              `json:"created_at"`
}

func CreateService(ctx context.Context, s *Service) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO services (
			active, user_id, category_id, subcategory_id,
			division_id, district_id, subdistrict_id,
			area, title, caption, description, price,
			features, hours, days,
			page_name, page_link, messenger_name, messenger_link
		) VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,
			$13,$14,$15,$16,$17,$18,$19
		)
	`, s.Active, s.UserID, s.CategoryID, s.SubcategoryID,
		s.DivisionID, s.DistrictID, s.SubdistrictID,
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
			id, active, user_id, category_id, subcategory_id,
			division_id, district_id, subdistrict_id,
			area, title, caption, description, price,
			features, hours, days,
			page_name, page_link, messenger_name, messenger_link,
			created_at
		FROM services
		WHERE id=$1
	`, id).Scan(
		&s.ID, &s.Active, &s.UserID, &s.CategoryID, &s.SubcategoryID,
		&s.DivisionID, &s.DistrictID, &s.SubdistrictID,
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

func UpdateService(ctx context.Context, s *Service) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE services
		SET active=$1, category_id=$2, subcategory_id=$3,
		    division_id=$4, district_id=$5, subdistrict_id=$6,
		    area=$7, title=$8, caption=$9, description=$10,
		    price=$11, features=$12, hours=$13, days=$14,
		    page_name=$15, page_link=$16, messenger_name=$17, messenger_link=$18
		WHERE id=$19
	`, s.Active, s.CategoryID, s.SubcategoryID,
		s.DivisionID, s.DistrictID, s.SubdistrictID,
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
		SELECT id, active, user_id, category_id, subcategory_id,
		       division_id, district_id, subdistrict_id,
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
			&s.ID, &s.Active, &s.UserID, &s.CategoryID, &s.SubcategoryID,
			&s.DivisionID, &s.DistrictID, &s.SubdistrictID,
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

func GetServicesByFilters(ctx context.Context, divisionID, districtID, subdistrictID int, categoryID, subcategoryID int64) ([]*Service, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT 
			id, active, user_id, category_id, subcategory_id,
			division_id, district_id, subdistrict_id,
			area, title, caption, description, price,
			features, hours, days,
			page_name, page_link, messenger_name, messenger_link,
			created_at
		FROM services
		WHERE division_id=$1 AND district_id=$2 AND subdistrict_id=$3
		  AND category_id=$4 AND subcategory_id=$5
		  AND active=TRUE
		ORDER BY created_at DESC
	`, divisionID, districtID, subdistrictID, categoryID, subcategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []*Service
	for rows.Next() {
		s := &Service{}
		if err := rows.Scan(
			&s.ID, &s.Active, &s.UserID, &s.CategoryID, &s.SubcategoryID,
			&s.DivisionID, &s.DistrictID, &s.SubdistrictID,
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
