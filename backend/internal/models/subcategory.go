package models

import (
	"backend/internal/db"
	"context"
	"time"
)

type SubCategory struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitzero"`
}

func CreateSubCategory(ctx context.Context, sc *SubCategory) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO sub_categories (category_id, name, description)
		VALUES ($1, $2, $3)
	`, sc.CategoryID, sc.Name, sc.Description)
	return err
}

func GetSubCategoryByID(ctx context.Context, id int64) (*SubCategory, error) {
	sc := &SubCategory{}
	err := db.Pool.QueryRow(ctx, `
		SELECT id, category_id, name, description, created_at
		FROM sub_categories
		WHERE id=$1
	`, id).Scan(&sc.ID, &sc.CategoryID, &sc.Name, &sc.Description, &sc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func GetAllSubCategories(ctx context.Context) ([]*SubCategory, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT id, category_id, name, description, created_at
		FROM sub_categories
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subCategories []*SubCategory
	for rows.Next() {
		sc := &SubCategory{}
		if err := rows.Scan(&sc.ID, &sc.CategoryID, &sc.Name, &sc.Description, &sc.CreatedAt); err != nil {
			return nil, err
		}
		subCategories = append(subCategories, sc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subCategories, nil
}

func UpdateSubCategory(ctx context.Context, sc *SubCategory) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE sub_categories
		SET category_id = $1,
		    name = $2,
		    description = $3
		WHERE id = $4
	`, sc.CategoryID, sc.Name, sc.Description, sc.ID)

	return err
}

func DeleteSubCategory(ctx context.Context, id int64) error {
	_, err := db.Pool.Exec(ctx, `
		DELETE FROM sub_categories
		WHERE id=$1
	`, id)
	return err
}
