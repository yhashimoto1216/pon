package model

import "time"

type Sentence struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"index"`
	Content   string `gorm:"type:text;not null"`
	Note      string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
