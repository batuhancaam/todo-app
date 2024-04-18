package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Task      string
	StartTime time.Time
	EndTime   time.Time
	Completed bool `gorm:"default:false"`
	UserID    uint
}
