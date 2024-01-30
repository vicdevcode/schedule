package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WeeklyEvent struct {
	ID        uuid.UUID `json:"id" gorm:"uuid;default:gen_random_uuid()"`
	Name      string    `json:"name"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	WeekDay   string    `json:"week_day"`
	WeekType  string    `json:"week_type"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
