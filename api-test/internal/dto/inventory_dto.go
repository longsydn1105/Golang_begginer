package dto

type InventoryCSVRow struct {
	ProductSKU    string `json:"product_sku"`
	CategoryCode  string `json:"category_code"`
	WarehouseCode string `json:"warehouse_code"`
	Quantity      int    `json:"quantity"`
	Type          string `json:"transaction_type"`
}

type ImportErrorDetail struct {
	Row     int         `json:"row"`
	Field   string      `json:"field"`
	Value   interface{} `json:"value"`
	Message string      `json:"message"`
}

type ImportResult struct {
	TotalRows   int                 `json:"total_rows"`   
	SuccessRows int                 `json:"success_rows"` 
	ErrorRows   int                 `json:"error_rows"`   
	ErrorList   []ImportErrorDetail `json:"error_list"`  
}
