package config

import "schule/internal/db/models"


func Seeds() []any {
	return []any{
		&models.Role{ID: 1, Name: "admin"},
		&models.Role{ID: 2, Name: "user"},
	}
}