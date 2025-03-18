package models

import (
	"fmt"

	"github.com/ShohruzNuraddinov/go-menu-bot/config"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	CategoryID  int
	IsActive    bool
}

func (p *Product) GetProductsByCategory(categoryID string) ([]Product, error) {
	db := config.GetDB()
	query := `SELECT id, name, description, price, category_id, is_active FROM products WHERE category_id = $1 AND is_active = true`

	rows, err := db.Query(query, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by category: %w", err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.IsActive)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, product)
	}

	// Check for errors after iterating
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through rows: %w", err)
	}

	return products, nil
}

func (p *Product) GetProductByID(productID string) (*Product, error) {
	db := config.GetDB()
	query := `SELECT id, name, description, price, category_id, is_active FROM products WHERE id = $1`

	err := db.QueryRow(query, productID).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CategoryID, &p.IsActive)
	if err != nil {
		return nil, fmt.Errorf("failed to get product by id: %w", err)
	}
	return p, nil
}
