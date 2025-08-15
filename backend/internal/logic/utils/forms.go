package utils

import (
	"encoding/json"
	"errors"
	"schule/internal/api/model"
)

func ValidateForm(form *model.Form) error {
	if form.Title == "" {
		return errors.New("title is required")
	}
	if form.Description == "" {
		return errors.New("description is required")
	}
	if form.Body == nil {
		return errors.New("body is required")
	}
	if form.FormEditors.Groups == nil && form.FormEditors.Users == nil {
		return errors.New("at least one form editor is required")
	}
	if form.FormViewer.Groups == nil && form.FormViewer.Users == nil {
		return errors.New("at least one form viewer is required")
	}
	// TODO implement body validation
	return nil
}

func InterfaceToBytes(data interface{}) []byte {
    if b, ok := data.([]byte); ok {
        return b
    }
    b, _ := json.Marshal(data)
    return b
}