package models

import (
	"gorm.io/gorm"
)

type Base struct {
	gorm.Model

	Description *string `gorm:"type:text"`
	Status      *string `gorm:"type:varchar(255);default:active"`
}
