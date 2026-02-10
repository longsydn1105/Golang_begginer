package model

import "github.com/google/uuid"

type Warehouse struct {
	ID   uuid.UUID `gorm:"uuid;default:gen_random;primaryKey"`
	Code string    `gorm:"unique;not null"`
}
