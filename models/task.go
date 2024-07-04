package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID    uint `json:"passportNumber"`
	TaskName  string
	StartTime time.Time
	EndTime   *time.Time
}
