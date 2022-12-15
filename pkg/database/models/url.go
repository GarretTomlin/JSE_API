package models

import (
	"time"
	"gorm.io/gorm"
)

func DbModels() {
	type UrlRecords struct {
		gorm.Model
		Url string
		Date time.Time
	}
}
