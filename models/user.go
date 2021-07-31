package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"type:VARCHAR(60);NOT NULL"`
	LastName  string `json:"last_name" gorm:"type:VARCHAR(60);NOT NULL"`
	Password  string `json:"password" gorm:"type:VARCHAR(255);NOT NULL"`
	Email     string `json:"email" gorm:"type:VARCHAR(254);uniqueIndex;NOT NULL"`
}
