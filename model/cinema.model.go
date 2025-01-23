package model

import "time"

type Cinema struct {
	ID        int       `gorm:"primaryKey"`
	Name      string    `gorm:"size:100"`
	Location  string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
