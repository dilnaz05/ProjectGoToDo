package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
	UserID   uint   `json:"user_id"`
}
