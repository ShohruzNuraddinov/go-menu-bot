package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          int `gorm:"primaryKey"`
	Name        string `gorm:"size:255"`
	Description string `gorm:"size:1020"`
	Price       float64 `gorm:"type:decimal(10,2)"`
	CategoryID  int
	IsActive    bool
}