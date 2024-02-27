package models

import (
	"github.com/jinzhu/gorm"
	"github.com/msrpr20/pkg/project/config"
)

var db *gorm.DB

type Book struct {
	gorm.model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.Automigrate(&Book{})
}
