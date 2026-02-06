package interfaces

import (
	"api-test/internal/dto"
	"mime/multipart"
)

type InventoryServiceInterface interface {
	ProccessUploadCSV(file multipart.File) (*dto.ImportResult, error)
}
