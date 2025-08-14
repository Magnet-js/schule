package service

import (
	"schule/internal/api/model"
	"schule/internal/db/models"
	"schule/internal/logic/utils"

	"github.com/google/uuid"
)

type FormsService struct {
	repo FormsRepository
}

type FormsRepository interface {
	Create(form *models.Form) error
	Update(form *models.Form) error
}

func NewFormsService(repo FormsRepository) *FormsService {
	return &FormsService{
		repo: repo,
	}
}

func (s *FormsService) CreateForm(form *model.Form) error {
	if err := utils.ValidateForm(form); err != nil {
		return err
	}

	// TODO create shared mapper 
	dbForm := &models.Form{
		ID: 			  uuid.New(),
		Title:             form.Title,
		Description:       form.Description,
		Body:              utils.InterfaceToBytes(form.Body),
		MaxSubmitsPerUser: form.MaxSubmitsPerUser,
	}

	return s.repo.Create(dbForm)
}

func (s *FormsService) UpdateForm(form *model.Form) error {
	if err := utils.ValidateForm(form); err != nil {
		return err
	}

	// TODO create shared mapper 
	dbForm := &models.Form{
		Title:             form.Title,
		Description:       form.Description,
		Body:              utils.InterfaceToBytes(form.Body),
		MaxSubmitsPerUser: form.MaxSubmitsPerUser,
	}

	return s.repo.Update(dbForm)
}