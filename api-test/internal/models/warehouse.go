package models

type Warehouse struct {
	ID   int    `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
}
