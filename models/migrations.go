package models

import (
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Category{}, &Product{}, &TelegramUser{})
	if err != nil {
		return err
	}
	return nil
}