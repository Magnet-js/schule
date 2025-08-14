package api

import (
	api_forms "schule/internal/api/forms"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	version := api.Group("/v1")

	version.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	api_forms.RegisterRoutes(version, db)
}
