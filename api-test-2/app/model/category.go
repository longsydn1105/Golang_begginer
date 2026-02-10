package model

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid;primaryKey"`
	Code string    `gorm:"unique; not null"`
	Name string
}
