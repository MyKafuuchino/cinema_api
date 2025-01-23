package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	FullName  string    `gorm:"size:100"`
	Email     string    `gorm:"size:100;unique"`
	Password  string    `gorm:"size:255"`
	Role      string    `gorm:"type:enum('USER','ADMIN');default:'USER'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
