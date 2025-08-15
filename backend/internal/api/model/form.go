package model

import "encoding/json"

type Form struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        json.RawMessage `json:"body"`

	MultiViewable bool   `json:"multi_viewable"`
	ApproveNeeded bool   `json:"approve_needed"`
	FormEditors   FormRole `json:"form_editors"`
	FormViewer    FormRole `json:"form_viewer"`
	FormApprovers FormRole `json:"form_approvers"`
}
type FormRole struct {
    Groups []string `json:"groups,omitempty"`
    Users  []string `json:"users,omitempty"`
}