package main

import (
	"api-test/cmd"
	"log"
)

func main() {
	server, err := cmd.NewServer("config.yaml")
	if err != nil {
		log.Fatalf("❌ Failed to init server: %v", err)
	}

	r := server.SetupRoutes()
	log.Println("🚀 Server running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
