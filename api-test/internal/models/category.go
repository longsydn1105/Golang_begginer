package models

type Category struct {
	ID   int    `gorm:"primaryKey"`
	Code string `gorm:"unique;not null"`
	Name string
}
