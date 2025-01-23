package model

import "time"

type Audit struct {
	ID        int       `gorm:"primaryKey"`
	Action    string    `gorm:"size:255"`
	UserID    int       `gorm:"index"`
	Timestamp time.Time `gorm:"type:timestamp"`
	Details   string    `gorm:"type:json"`
}
