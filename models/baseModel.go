package models

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"PrimarKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
