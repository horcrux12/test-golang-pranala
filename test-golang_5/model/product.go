package model

import "gorm.io/gorm"

type Product struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Qty       int    `json:"qty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}