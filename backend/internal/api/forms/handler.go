package api_forms

import (
	"net/http"
	"schule/internal/db/repository"
	"schule/internal/logic/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
	formsService *service.FormsService
}

func Newhandler(db *gorm.DB) *handler {
	return &handler{
		db: db,
		formsService: service.NewFormsService(repository.NewFormsRepository(db)),
	}
}

func (h *handler) CreateForm(c *gin.Context) {
	if err := h.handleCreateForm(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Form created successfully"})
}
