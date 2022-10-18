package models

import (
	"time"
)

type Model struct {
	ID        uint `gorm:"PrimaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
