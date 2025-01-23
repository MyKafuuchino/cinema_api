package model

import "time"

type Ticket struct {
	ID          int       `gorm:"primaryKey"`
	UserID      int       `gorm:"index"`
	ScreeningID int       `gorm:"index"`
	SeatNumber  string    `gorm:"size:10"`
	Status      string    `gorm:"type:enum('booked','paid')"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
