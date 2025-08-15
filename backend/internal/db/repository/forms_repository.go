package repository

import (
	"schule/internal/db/models"

	"gorm.io/gorm"
)

type FormsRepository struct {
	db *gorm.DB
}

func NewFormsRepository(db *gorm.DB) *FormsRepository {
	return &FormsRepository{
		db: db,
	}
}

func (r *FormsRepository) Create(form *models.Form) error {
	return r.db.Create(form).Error
}

func (r *FormsRepository) Update(form *models.Form) error {
	return r.db.Save(form).Error
}