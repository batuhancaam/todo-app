package model

import "time"

type Todo struct {
	ID        uint `gorm:"primaryKey"`
	Task      string
	StartTime time.Time
	EndTime   time.Time
	Completed bool `gorm:"default:false"`
}
