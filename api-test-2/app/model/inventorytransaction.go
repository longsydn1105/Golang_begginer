package model

import (
	"time"

	"github.com/google/uuid"
)

type InventoryTransaction struct {
	ID              uuid.UUID `gorm:"type:uuid;defautl:gen_random;primaryKey"`
	ProductID       uuid.UUID `gorm:"type:uuid;default:gen_random"`
	CategoryID      uuid.UUID `gorm:"type:uuid;default:gen_random"`
	WarehouseID     uuid.UUID `gorm:"type:uuid;default:gen_random"`
	Quantity        int       `gorm:"not null"`
	TransactionType string    `gorm:"size:10;not null"`
	CreateAt        time.Time `gorm:"autoCreateTime"`
}
