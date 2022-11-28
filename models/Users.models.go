package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Firstname string  `gorm:"not null"`
	Lastname  string  `gorm:"not null"`
	Email     string  `gorm:"not null;unique_index"`
	Task      []Tasks `gorm:"foreignKey:UserID"`
}
