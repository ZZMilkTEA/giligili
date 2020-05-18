package model

import (
	"github.com/jinzhu/gorm"
)

type ReviewLog struct {
	gorm.Model
	MediaId        uint   `gorm:"not null"`
	MediaType      string `gorm:"type:enum('video','audio');not null"`
	ReviewerId     uint   `gorm:"not null"`
	StatusForward  string `gorm:"type:enum('pending_review','passed','not_passed'); default:'pending_review'"`
	StatusBackward string `gorm:"type:enum('pending_review','passed','not_passed'); default:'pending_review'"`
	Remark         string
}
