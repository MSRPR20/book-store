package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	db = d
}
func getDB() *gorm.DB {
	return db
}