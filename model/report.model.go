package model

import "time"

type Report struct {
	ID         int       `gorm:"primaryKey"`
	ReportName string    `gorm:"size:100"`
	Content    string    `gorm:"type:json"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
