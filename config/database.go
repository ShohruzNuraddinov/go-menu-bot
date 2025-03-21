package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	psqlInfo := GetConfig().psqlInfo

	var err error
	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")
	return DB
}
