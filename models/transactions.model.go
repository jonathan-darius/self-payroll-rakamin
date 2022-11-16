package models

import "gorm.io/gorm"

type Transaction struct {
	Amount int    `json:"amount"`
	Note   string `json:"note"`
	Type   string `json:"type"`
	gorm.Model
}
