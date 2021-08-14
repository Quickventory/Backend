package models

import (
	"time"

	"gorm.io/gorm"
)

type AccessToken struct {
	gorm.Model
	UserID    uint
	User      User
	IssuedAt  time.Time
	Revoked   bool
	ExpiresAt time.Time
}
