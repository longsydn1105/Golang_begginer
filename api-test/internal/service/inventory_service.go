package service

import (
	"api-test/internal/dto"
	"api-test/internal/interfaces"
	"api-test/internal/models"
	"encoding/csv"
	"errors"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
)

type inventoryService struct {
	repo interfaces.InventoryRepositoryInterface
}

func NewInventoryService(repo interfaces.InventoryRepositoryInterface) interfaces.InventoryServiceInterface {
	return &inventoryService{repo: repo}
}

func (s *inventoryService) ProccessUploadCSV(file multipart.File) (*dto.ImportResult, error) {
	productMap, categoryMap, warehouseMap, err := s.repo.GetMasterData()
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	if err := s.validateHeader(reader); err != nil {
		return nil, err
	}

	var validRecords []models.InventoryTransaction
	var errorList []dto.ImportErrorDetail
	rowNumber := 1

	// 4. Loop đọc file
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errorList = append(errorList, dto.ImportErrorDetail{Row: rowNumber + 1, Message: "Cannot read row data"})
			rowNumber++
			continue
		}
		rowNumber++

		// --- ĐOẠN NÀY ĐÃ ĐƯỢC TÁCH RA HÀM RIÊNG ---
		// Gọi đệ tử `validateAndMapRow` vào làm việc chi tiết
		transaction, errDetail := s.validateAndMapRow(record, rowNumber, productMap, categoryMap, warehouseMap)

		if errDetail != nil {
			// Nếu có lỗi -> Ghi sổ đen
			errorList = append(errorList, *errDetail)
		} else {
			// Nếu ngon lành -> Gom hàng
			validRecords = append(validRecords, *transaction)
		}
	}

	// 5. Bulk Insert
	if len(validRecords) > 0 {
		if err := s.repo.BulkInsert(validRecords); err != nil {
			return nil, err
		}
	}

	// 6. Trả kết quả
	return &dto.ImportResult{
		TotalRows:   rowNumber - 1,
		SuccessRows: len(validRecords),
		ErrorRows:   len(errorList),
		ErrorList:   errorList,
	}, nil
}

func (s *inventoryService) validateHeader(reader *csv.Reader) error {
	header, err := reader.Read()
	if err != nil {
		return errors.New("cannot read CSV header")
	}

	expectedHeader := []string{"product_sku", "category_code", "warehouse_code", "quantity", "transaction_type"}
	if len(header) != len(expectedHeader) {
		return errors.New("invalid CSV format: incorrect number of columns")
	}
	for i, col := range expectedHeader {
		if strings.TrimSpace(header[i]) != col {
			return errors.New("invalid CSV header: col " + strconv.Itoa(i+1) + " must be " + col)
		}
	}
	return nil
}

func (s *inventoryService) validateAndMapRow(
	record []string,
	rowNum int,
	pMap, cMap, wMap map[string]int,
) (*models.InventoryTransaction, *dto.ImportErrorDetail) {
	sku := strings.TrimSpace(record[0])
	catCode := strings.TrimSpace(record[1])
	whCode := strings.TrimSpace(record[2])
	qtyStr := strings.TrimSpace(record[3])
	txType := strings.TrimSpace(record[4])

	prodID, ok := pMap[sku]
	if !ok {
		return nil, &dto.ImportErrorDetail{Row: rowNum, Field: "product_sku", Value: sku, Message: "product not found"}
	}

	catID, ok := cMap[catCode]
	if !ok {
		return nil, &dto.ImportErrorDetail{Row: rowNum, Field: "category_code", Value: catCode, Message: "category not found"}
	}

	whID, ok := wMap[whCode]
	if !ok {
		return nil, &dto.ImportErrorDetail{Row: rowNum, Field: "warehouse_code", Value: whCode, Message: "warehouse not found"}
	}

	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		return nil, &dto.ImportErrorDetail{Row: rowNum, Field: "quantity", Value: qtyStr, Message: "quantity must be a number"}
	}

	if txType == "IN" && qty < 0 {
		return nil, &dto.ImportErrorDetail{Row: rowNum, Field: "quantity", Value: qty, Message: "quantity must be >= 0 for IN transaction"}
	}

	if txType != "IN" && txType != "OUT" {
		return nil, &dto.ImportErrorDetail{Row: rowNum, Field: "transaction_type", Value: txType, Message: "must be IN or OUT"}
	}

	return &models.InventoryTransaction{
		ProductID:       prodID,
		CategoryID:      catID,
		WarehouseID:     whID,
		Quantity:        qty,
		TransactionType: txType,
	}, nil
}
