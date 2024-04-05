package app

import (
	"test_golang_5/model"

	"gorm.io/gorm"
)

type appAttribute struct {
	DB *gorm.DB
}

var AppAtrribute = appAttribute{}

func InitAppAttribute() {
	AppAtrribute.DB = DBConnect()

	AppAtrribute.DB.AutoMigrate(model.Product{})
}