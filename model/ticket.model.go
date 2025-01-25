package model

import "time"

type Ticket struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"index"`
	ScreeningID uint      `gorm:"index"`
	SeatNumber  int       `gorm:""`
	Status      string    `gorm:"type:enum('booked','paid', 'canceled');default:'booked'"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
