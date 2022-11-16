package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	SecretID   string    `json:"secret_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	PositionID int       `json:"position_id"`
	Position   *Position `json:"position" gorm:"foreignKey:id"`
}
