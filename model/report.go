package model

import "github.com/jinzhu/gorm"

type Report struct {
	gorm.Model
	ReporterId  uint
	ReportType  string `gorm:"type:enum('video','audio','user','comment');not null"`
	ReportedId  uint   `gorm:"not null"`
	Description string
}
