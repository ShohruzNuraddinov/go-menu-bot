package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID       int       `gorm:"primaryKey"`
	Name     string    `gorm:"size:255"`
	IsActive bool      `gorm:"default:true"`
	Products []Product `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
