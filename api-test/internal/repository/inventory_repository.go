package repository

import (
	"api-test/internal/interfaces"
	"api-test/internal/models"

	"gorm.io/gorm"
)

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) interfaces.InventoryRepositoryInterface {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) GetMasterData() (map[string]int, map[string]int, map[string]int, error) {
	productMap := make(map[string]int)
	categoryMap := make(map[string]int)
	warehouseMap := make(map[string]int)

	var products []models.Product
	var categories []models.Category
	var warehouses []models.Warehouse

	if err := r.db.Select("id, sku").Find(&products).Error; err != nil {
		return nil, nil, nil, err
	}
	if err := r.db.Select("id, code").Find(&categories).Error; err != nil {
		return nil, nil, nil, err
	}
	if err := r.db.Select("id, code").Find(&warehouses).Error; err != nil {
		return nil, nil, nil, err
	}

	for _, p := range products {
		productMap[p.SKU] = p.ID
	}
	for _, p := range categories {
		categoryMap[p.Code] = p.ID
	}
	for _, p := range warehouses {
		warehouseMap[p.Code] = p.ID
	}
	return productMap, categoryMap, warehouseMap, nil
}

func (r *inventoryRepository) BulkInsert(transactions []models.InventoryTransaction) error {
	batchSize := 1000
	return r.db.CreateInBatches(transactions, batchSize).Error
}
