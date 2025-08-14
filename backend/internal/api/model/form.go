package model

type Form struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        interface{} `json:"body"`
	MaxSubmitsPerUser int    `json:"max_submits_per_user"`
}