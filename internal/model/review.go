package model

import "time"

type Review struct {
	ID           uint `gorm:"primaryKey"`
	SentenceID   uint
	UserID       uint
	Repetition   uint
	Interval     uint
	Easiness     float64
	NextReview   time.Time
	LastReviewed time.Time
}
