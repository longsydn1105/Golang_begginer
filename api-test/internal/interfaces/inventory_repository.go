package interfaces

import (
	"api-test/internal/models"
)

type InventoryRepositoryInterface interface {
	//Lấy dữ liệu từ master (Product, Category, Warehouse) lên RAM
	// Trả về 3 cái Map để tra cứu cho nhanh: map[Code]ID
	GetMasterData() (map[string]int, map[string]int, map[string]int, error)
	BulkInsert(transactions []models.InventoryTransaction) error
}
