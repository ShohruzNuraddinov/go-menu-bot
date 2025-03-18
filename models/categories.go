package models

import (
	"fmt"

	"github.com/ShohruzNuraddinov/go-menu-bot/config"
)

type Category struct {
	ID       int
	Name     string
	IsActive bool
}

func (c *Category) GetCategories() ([]Category, error) {
	db := config.GetDB()
	query := `SELECT * FROM category WHERE is_active = true`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.Name, &category.IsActive)
		if err != nil {
			return nil, fmt.Errorf("failed to scan category: %w", err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) GetCategoryByID(categoryID string) (*Category, error) {
	db := config.GetDB()
	query := `SELECT * FROM category WHERE id = $1`

	err := db.QueryRow(query, categoryID).Scan(&c.ID, &c.Name, &c.IsActive)
	if err != nil {
		return nil, fmt.Errorf("failed to get category by id: %w", err)
	}
	return c, nil
}