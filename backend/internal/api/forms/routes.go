package api_forms

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	formsGroup := r.Group("/forms")
	handler := Newhandler(db)

	formsGroup.PUT("/create", handler.CreateForm)

}