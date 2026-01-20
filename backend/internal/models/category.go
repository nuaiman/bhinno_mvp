package models

import (
	"backend/internal/db"
	"context"
	"time"
)

type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitzero"`
}

func CreateCategory(ctx context.Context, c *Category) error {
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
	`, c.Name, c.Description)

	return err
}

func GetCategoryByID(ctx context.Context, id int64) (*Category, error) {
	c := &Category{}
	err := db.Pool.QueryRow(ctx, `
		SELECT id, name, description, created_at
		FROM categories
		WHERE id = $1
	`, id).Scan(
		&c.ID,
		&c.Name,
		&c.Description,
		&c.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetAllCategories(ctx context.Context) ([]*Category, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT id, name, description, created_at
		FROM categories
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*Category
	for rows.Next() {
		c := &Category{}
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
			&c.CreatedAt,
		); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func UpdateCategory(ctx context.Context, c *Category) error {
	_, err := db.Pool.Exec(ctx, `
		UPDATE categories
		SET name = $1,
		    description = $2
		WHERE id = $3
	`, c.Name, c.Description, c.ID)

	return err
}

func DeleteCategory(ctx context.Context, id int64) error {
	_, err := db.Pool.Exec(ctx, `
		DELETE FROM categories
		WHERE id = $1
	`, id)

	return err
}
