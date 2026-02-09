package handler

import (
	"api-test/internal/dto"
	"api-test/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

const MaxSizeFile = 1000 << 20

type InventoryHandler struct {
	service interfaces.InventoryServiceInterface
}

func NewInventoryHandler(service interfaces.InventoryServiceInterface) *InventoryHandler {
	return &InventoryHandler{service: service}
}

func (h *InventoryHandler) UploadCSV(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxSizeFile)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse(400, "File is required", err.Error()))
		return
	}

	if fileHeader.Size > MaxSizeFile {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse(400, "File too large", "Max size is 50MB"))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse(500, "Cannot open file", err.Error()))
		return
	}
	defer file.Close()

	result, err := h.service.ProccessUploadCSV(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse(500, "Processing failed", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResponse(result))
}
