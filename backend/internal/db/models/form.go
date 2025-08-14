package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Form struct {
	ID                uuid.UUID       `json:"id" gorm:"primaryKey"`
	Title             string          `json:"title"`
	Description       string          `json:"description"`
	MaxSubmitsPerUser int             `json:"max_submits_per_user"`
	Body              json.RawMessage `json:"body"`

	// CreatorID
	// View logic who is able to view the form
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
