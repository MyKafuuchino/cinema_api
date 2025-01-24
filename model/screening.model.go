package model

import "time"

type Screening struct {
	ID            uint      `gorm:"primaryKey"`
	MovieID       uint      `gorm:"index"`
	CinemaID      uint      `gorm:"index"`
	ScreeningTime time.Time `gorm:"type:timestamp"`
	Price         float64   `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
