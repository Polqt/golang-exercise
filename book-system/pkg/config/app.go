package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db * gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "host=localhost user=postgres password=hidalgo001 dbname=bookstore port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}