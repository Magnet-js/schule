package config

import (
	"encoding/json"
	"schule/internal/db/models"
	"time"

	"github.com/google/uuid"
)


func Seeds() []any {
	return []any{
		&models.Role{ID: 1, Name: "admin"},
		&models.Role{ID: 2, Name: "user"},
		// add sample form data and for submissions that are working fine and are best practise setuped
		&models.Form{ID: uuid.New(), Title: "Sample Form", Description: "This is a sample form", Body: json.RawMessage(`{"questions": [{"id": "q1", "type": "text", "label": "What is your name?"}, {"id": "q2", "type": "text", "label": "What is your email?"}]}`), MultiViewable: true, ApproveNeeded: false, FormEditors: models.FormRole{Groups: []string{"editors"}, Users: []string{"user1"}}, FormViewer: models.FormRole{Groups: []string{"viewers"}, Users: []string{"user2"}}, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		&models.Submission{ID: uuid.New(), FormID: uuid.New(), Answer: json.RawMessage(`{"q1": "answer1", "q2": "answer2"}`), CreatedAt: time.Now()},
	}
}