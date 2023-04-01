package models

import "time"

type Book struct {
	ID        uint   `gorm:"primaryKey"`
	NameBook  string `gorm:"not null;type:varchar(191)"`
	Author    string `gorm:"not null;type:varchar(191)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
