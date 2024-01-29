package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Example struct {
	gorm.Model
	ID      uuid.UUID `gorm:"uuid;default:gen_random_uuid()"`
	Message string
}
