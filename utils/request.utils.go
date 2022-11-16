package utils

import "github.com/jinzhu/gorm"

type TopUpUtils struct {
	Balance int `json:"balance" validate:"required"`
}

type UserRequest struct {
	gorm.Model
	SecretID   string `json:"secret_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	PositionID int    `json:"position_id"`
}

type WithdrawRequest struct {
	Id        int    `json:"id"`
	Secret_id string `json:"secret_id"`
}
