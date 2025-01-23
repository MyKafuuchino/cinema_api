package model

import "time"

type Screening struct {
	ID            int       `gorm:"primaryKey"`
	MovieID       int       `gorm:"index"`
	CinemaID      int       `gorm:"index"`
	ScreeningTime time.Time `gorm:"type:timestamp"`
	Price         float64   `gorm:"type:decimal(10,2)"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
