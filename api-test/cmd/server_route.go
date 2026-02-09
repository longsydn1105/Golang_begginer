package cmd

import "github.com/gin-gonic/gin"

func (s *Server) SetupRoutes() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/upload", s.InventoryHandler.UploadCSV)
	}
	return r
}
