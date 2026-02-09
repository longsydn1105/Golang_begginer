package models

import "time"

type InventoryTransaction struct {
	ID              int64     `gorm:"primaryKey;autoIncrement"`
	ProductID       int       `gorm:"not null"`
	CategoryID      int       `gorm:"not null"`
	WarehouseID     int       `gorm:"not null"`
	Quantity        int       `gorm:"not null"`
	TransactionType string    `gorm:"size:10;not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}

func (InventoryTransaction) TableName() string {
	return "inventory_transactions"
}
