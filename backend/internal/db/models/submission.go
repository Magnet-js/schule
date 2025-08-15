package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Submission struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey"`
	// UserID
	FormID uuid.UUID `json:"form_id"`
	Answer json.RawMessage `json:"answer"`

	CreatedAt time.Time `json:"created_at"`
}