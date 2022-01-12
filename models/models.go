package models

import "gorm.io/gorm"

// Quotes model
type Quotes struct {
	gorm.Model

	Id     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}
