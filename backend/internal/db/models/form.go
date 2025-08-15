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
	Body              json.RawMessage `json:"body" gorm:"type:jsonb"`
	MultiViewable     bool            `json:"multi_viewable"`
	ApproveNeeded     bool            `json:"approve_needed"`
	FormEditors       FormRole       `json:"form_editors" gorm:"type:jsonb"`
	FormViewer        FormRole       `json:"form_viewer" gorm:"type:jsonb"`
	FormApprovers     FormRole       `json:"form_approvers" gorm:"type:jsonb"`

	// CreatorID
	// View logic who is able to view the form
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type FormRole struct {
	Groups []string `json:"groups,omitempty"`
	Users  []string `json:"users,omitempty"`
}