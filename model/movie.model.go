package model

import "time"

type Movie struct {
	ID          int       `gorm:"primaryKey"`
	Title       string    `gorm:"size:100"`
	Description string    `gorm:"type:text"`
	Genre       string    `gorm:"size:50"`
	Duration    int       `gorm:"type:int"`
	ReleaseDate time.Time `gorm:"type:date"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
