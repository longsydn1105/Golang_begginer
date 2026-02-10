package model

import "github.com/google/uuid"

type Product struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random;primaryKey"`
	SKU        string    `gorm:"unique;not null"`
	CategoryId int
	Category   Category `gorm:"foreignKey:CategoryId"`
}
