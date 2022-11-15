package models

import "github.com/jinzhu/gorm"

type Position struct {
	Name   string `json:"name" validate:"required"`
	Salary int    `json:"salary" validate:"required"`
	gorm.Model
}
