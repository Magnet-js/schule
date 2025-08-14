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
	if form.MaxSubmitsPerUser <= 0 {
		return errors.New("max submits per user must be positive")
	}
	if form.Body == nil {
		return errors.New("body is required")
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