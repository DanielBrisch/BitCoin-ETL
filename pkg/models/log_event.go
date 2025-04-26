package models

import "time"

type LogEvent struct {
	ID        uint      `gorm:"primaryKey"`
	Level     string    `gorm:"column:level"`
	Message   string    `gorm:"column:message"`
	Timestamp time.Time `gorm:"column:timestamp"`
}
