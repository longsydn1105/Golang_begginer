package service

import (
	"api-test/internal/dto"
	"api-test/internal/interfaces"
	"api-test/internal/models"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	WorkerCount = 10
	BatchSize   = 5000
)

type inventoryService struct {
	repo interfaces.InventoryRepositoryInterface
}

func NewInventoryService(repo interfaces.InventoryRepositoryInterface) interfaces.InventoryServiceInterface {
	return &inventoryService{repo: repo}
}

func (s *inventoryService) ProccessUploadCSV(file multipart.File) (*dto.ImportResult, error) {
	startTotal := time.Now()

	productMap, categoryMap, warehouseMap, err := s.repo.GetMasterData()
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)

	if err := s.validateHeader(reader); err != nil {
		return nil, err
	}

	// --- THIẾT LẬP WORKER POOL ---
	jobs := make(chan []models.InventoryTransaction, WorkerCount)
	results := make(chan int, WorkerCount)
	var wg sync.WaitGroup

	// Khởi động các workers
	for w := 0; w < WorkerCount; w++ {
		wg.Add(1)
		go s.workerInsert(w, jobs, results, &wg)
	}

	totalSuccess := 0
	done := make(chan bool)
	go func() {
		for count := range results {
			totalSuccess += count
		}
		done <- true
	}()

	var currentBatch []models.InventoryTransaction
	var errorList []dto.ImportErrorDetail
	rowNumber := 1

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

		transaction, errDetail := s.validateAndMapRow(record, rowNumber, productMap, categoryMap, warehouseMap)

		if errDetail != nil {
			errorList = append(errorList, *errDetail)
		} else {
			currentBatch = append(currentBatch, *transaction)
		}

		if len(currentBatch) >= BatchSize {
			batchToSend := make([]models.InventoryTransaction, len(currentBatch))
			copy(batchToSend, currentBatch)

			jobs <- batchToSend
			currentBatch = make([]models.InventoryTransaction, 0, BatchSize)
		}
	}

	if len(currentBatch) > 0 {
		jobs <- currentBatch
	}

	close(jobs)
	wg.Wait()
	close(results)
	<-done

	totalDuration := time.Since(startTotal)
	fmt.Println("------------------------------------------------")
	fmt.Printf("🚀 Tốc độ BÀN THỜ (Validation + Concurrency)\n")
	fmt.Printf("✅ Tổng thời gian: %v\n", totalDuration)
	fmt.Printf("📦 Đã Insert thành công: %d dòng\n", totalSuccess)
	fmt.Printf("❌ Số dòng lỗi: %d dòng\n", len(errorList))

	if totalDuration.Seconds() > 0 {

		speed := float64(rowNumber-1) / totalDuration.Seconds()
		fmt.Printf("⚡ Tốc độ xử lý: %.0f dòng/giây\n", speed)
	}
	fmt.Println("------------------------------------------------")

	return &dto.ImportResult{
		TotalRows:   rowNumber - 1,
		SuccessRows: totalSuccess,
		ErrorRows:   len(errorList),
		ErrorList:   errorList,
	}, nil
}

func (s *inventoryService) workerInsert(id int, jobs <-chan []models.InventoryTransaction, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for batch := range jobs {
		if err := s.repo.BulkInsert(batch); err == nil {
			results <- len(batch)
		} else {
			fmt.Printf("Worker %d bị lỗi: %v\n", id, err)
		}
	}
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
