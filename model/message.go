package model

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Content string `gorm:"type:text(150);not null"`
	Password  string `gorm:"size:255;not null"`
}