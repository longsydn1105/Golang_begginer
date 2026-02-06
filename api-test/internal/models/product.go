package models

type Product struct {
	ID         int    `gorm:"primaryKey"`
	SKU        string `gorm:"unique;not null"`
	CategoryId int
	Category   Category `gorm:"foreignKey:CategoryID"`
}
