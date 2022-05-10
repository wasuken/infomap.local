package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title       string
	Description string
	PubDate     time.Time
	Source      string
	Link        string
}
