package models

import "github.com/jinzhu/gorm"

type InDB struct {
	DB *gorm.DB
}
