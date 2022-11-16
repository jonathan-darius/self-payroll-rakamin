package models

import "github.com/jinzhu/gorm"

type Company struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Balance int    `json:"balance"`
}
