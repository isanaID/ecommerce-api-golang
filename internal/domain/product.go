package domain

import (
	"time"
)

type Product struct {
	ID          int64     `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Price       float64   `gorm:"size:255" json:"price"`
	Stock       int64     `gorm:"size:255" json:"stock"`
	CategoryID  int64     `gorm:"index" json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	ID          int64     `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"size:255" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	Products    []Product `gorm:"foreignkey:CategoryID" json:"products"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
