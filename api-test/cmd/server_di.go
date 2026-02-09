package cmd

import (
	"api-test/config"
	"api-test/internal/handler"
	"api-test/internal/repository"
	"api-test/internal/service"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	Config           *config.Config
	DB               *gorm.DB
	InventoryHandler *handler.InventoryHandler
}

func NewServer(configPath string) (*Server, error) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.Dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("✅ Connected to Database!")

	inventoryRepo := repository.NewInventoryRepository(db)
	inventoryService := service.NewInventoryService(inventoryRepo)
	inventoryHandler := handler.NewInventoryHandler(inventoryService)

	return &Server{
		Config:           cfg,
		DB:               db,
		InventoryHandler: inventoryHandler,
	}, nil
}
