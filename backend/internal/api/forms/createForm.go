package api_forms

import (
	"net/http"
	"schule/internal/api/model"

	"github.com/gin-gonic/gin"
)

func (h *handler) handleCreateForm(c *gin.Context) error {
	var form model.Form
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	if err := h.formsService.CreateForm(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	c.JSON(http.StatusCreated, form)
	return nil
}
